package queries

func GetComment(pid string) string {
	return "select c.UserId, u.Name, c.Comment, c.LastUpDate from Comment c, Users u where u.Id = c.UserId and c.ProductId = " + pid
}

func CheckCommentExist(uid, pid string) string {
	return "select exists(select Id from Comment where UserId = " + uid + " and ProductId = " + pid + ")"
}
