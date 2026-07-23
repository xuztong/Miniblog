package impl_test

import (
	"blog/apps/blog"
	"testing"
)

func TestCreate(t *testing.T) {
	req := blog.NewCreateBlogRequest()
	req.Title = "q1"
	req.Author = "q1"
	req.Content = "内容"
	req.Summary = "概要"
	req.Tags = map[string]string{"hello": "hellp"}
	ins, err := serviceImpl.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	ins, err := serviceImpl.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeBlog(t *testing.T) {
	req := blog.NewDesibleBlogRequest("1")
	ins, err := serviceImpl.DesibleBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestPathchUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("7")
	req.UpdateMode = blog.UPDATE_MODE_PATCH
	req.Title = "更新后文章标题"
	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestPutUpdateBlog(t *testing.T) {
	req := blog.NewUpdateBlogRequest("7")
	req.UpdateMode = blog.UPDATE_MODE_PUT
	req.Title = "更新后文章标题put"
	req.Author = "patch"
	req.Content = "path"
	ins, err := serviceImpl.UpdateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestUpdateBlogStatus(t *testing.T) {
	req := blog.NewUpdateStatusBlogRequest("7")
	req.SetStatus(blog.STATUS_PUBLISH)
	ins, err := serviceImpl.UpdateStatusBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	req := blog.NewDeleteBlogRequest("7")
	ins, err := serviceImpl.DeleteBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
