package impl

import (
	"blog/apps/blog"
	"blog/exception"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/redis/go-redis/v9"
)

func blogDetailKey(id string) string {
	return "blog:detail:" + id
}

func blogListKey(req *blog.QueryBlogRequest) string {
	return fmt.Sprintf("blog:list:%s:%d:%d", req.KeyWords, req.PageNumber, req.PageSize)
}

func (i *BlogServiceImpl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	//校验参数
	if err := in.Validate(); err != nil {
		return nil, err
	}
	//创建博客
	bl := blog.NewBlog()
	bl.CreateBlogRequest = in

	if err := i.db.WithContext(ctx).Create(bl).Error; err != nil {
		return nil, err
	}

	i.deleteBlogListCache(ctx)
	return bl, nil

}
func (i *BlogServiceImpl) DeleteBlog(ctx context.Context, in *blog.DeleteBlogRequest) (*blog.Blog, error) {
	ins, err := i.DesibleBlog(ctx, blog.NewDesibleBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	if err := i.db.WithContext(ctx).Where("id=?", in.BlogId).Delete(&blog.Blog{}).Error; err != nil {
		return nil, err
	}

	i.cache.Del(ctx, blogDetailKey(in.BlogId))

	return ins, err
}
func (i *BlogServiceImpl) UpdateBlog(ctx context.Context, in *blog.UpdateBlogRequest) (*blog.Blog, error) {
	ins, err := i.DesibleBlog(ctx, blog.NewDesibleBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	switch in.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		ins.CreateBlogRequest = in.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		err := mergo.MergeWithOverwrite(ins.CreateBlogRequest, in.CreateBlogRequest)
		if err != nil {
			return nil, err
		}
	}
	if err := ins.CreateBlogRequest.Validate(); err != nil {
		return nil, exception.ErrValidateFailed("%s errvalidate", err.Error())
	}

	if err := i.db.WithContext(ctx).Table("blogs").Save(ins).Error; err != nil {
		return nil, err
	}

	//mysql更新成功后，删除详细缓存
	i.cache.Del(ctx, blogDetailKey(in.BlogId))

	//清除列表缓存
	i.deleteBlogListCache(ctx)

	return ins, nil
}
func (i *BlogServiceImpl) QueryBlog(ctx context.Context, in *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	//1.先查redis
	cacheKey := blogListKey(in)
	data, err := i.cache.Get(ctx, cacheKey).Bytes()
	if err == nil {
		ins := blog.NewBlogSet()
		if jsonErr := json.Unmarshal(data, ins); jsonErr == nil {
			return ins, nil
		}
		i.cache.Del(ctx, cacheKey)
	}

	//2.查Mysql
	ins := blog.NewBlogSet()
	query := i.db.WithContext(ctx).Table("blogs")
	if in.KeyWords != "" {
		query = query.Where("title LIKE ?", "%"+in.KeyWords+"%")
	}
	if err := query.Count(&ins.Total).Error; err != nil {
		return nil, err
	}
	if err := query.Limit(in.PageSize).Offset(in.Offset()).Find(&ins.Items).Error; err != nil {
		return nil, err
	}

	//3.写入redis
	if cacheData, marshalErr := json.Marshal(ins); marshalErr == nil {
		i.cache.Set(ctx, cacheKey, cacheData, 1*time.Minute)
	}

	return ins, nil
}
func (i *BlogServiceImpl) DesibleBlog(ctx context.Context, in *blog.DesibleBlogRequest) (*blog.Blog, error) {
	//1.查redis
	cacheKey := blogDetailKey(in.BlogId)
	data, err := i.cache.Get(ctx, cacheKey).Bytes()
	if err == nil {
		ins := blog.NewBlog()
		if jsonErr := json.Unmarshal(data, ins); jsonErr == nil {
			return ins, nil
		}
		i.cache.Del(ctx, cacheKey)
	} else if err != redis.Nil {

	}

	//2.查数据库
	ins := blog.NewBlog()
	if err := i.db.WithContext(ctx).Where("id=?", in.BlogId).First(ins).Error; err != nil {
		return nil, err
	}

	//3.将数据库的数据存入redis
	if cacheData, marshalErr := json.Marshal(ins); marshalErr == nil {
		i.cache.Set(ctx, cacheKey, cacheData, 5*time.Minute)
	}
	return ins, nil
}
func (i *BlogServiceImpl) UpdateStatusBlog(ctx context.Context, in *blog.UpdateStatusBlogRequest) (*blog.Blog, error) {
	ins, err := i.DesibleBlog(ctx, blog.NewDesibleBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}
	ins.ChangeBlogStatusRequest = in.ChangeBlogStatusRequest
	ins.SetStatus(ins.Status)
	if err := i.db.WithContext(ctx).Table("blogs").Where("id=?", in.BlogId).Updates(ins.ChangeBlogStatusRequest).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *BlogServiceImpl) deleteBlogListCache(ctx context.Context) {
	iter := i.cache.Scan(ctx, 0, "blog:list:*", 100).Iterator()
	var keys []string

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return
	}

	if len(keys) > 0 {
		i.cache.Del(ctx, keys...)
	}

}

/*
func (i *BlogServiceImpl) UpdateBlog(
    ctx context.Context,
    in *blog.UpdateBlogRequest) (
    *blog.Blog, error) {

    // 1. 查询现有数据
    ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
    if err != nil {
        return nil, err
    }

    // 2. 统一使用 mergo 增量合并（不再判断 UpdateMode）
    err = mergo.MergeWithOverwrite(ins.CreateBlogRequest, in.CreateBlogRequest)
    if err != nil {
        return nil, err
    }

    // 3. 校验 & 保存
    if err := ins.CreateBlogRequest.Validate(); err != nil {
        return nil, exception.ErrValidateFailed(err.Error())
    }

    err = i.db.WithContext(ctx).Table("blogs").Save(ins).Error
    if err != nil {
        return nil, err
    }
    return ins, nil
}
*/
