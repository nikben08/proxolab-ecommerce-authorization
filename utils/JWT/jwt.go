func generateJWT(username string, userId int, accessLevel int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      userId,
		"username":    username,
		"accessLevel": accessLevel,
	})
	secretKey := []byte(config.Config("JWT_SECRET_FOR_LOCAL"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, err
}