package negocio

import (
	"cloud.google.com/go/spanner"
	"context"
	"github.com/jaimemr86/clases"
	"google.golang.org/api/option"
)

/******DESARROLLO*******/
var myDBPU = "projects/neodatabases/instances/neodata/databases/neodatapu2020desarrollo"
var PathCredencial = `{
  "type": "service_account",
  "project_id": "neodatabases",
  "private_key_id": "37f6dc1c58150fa5a2437644d47f4fe31f85a8f4",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDJJbuT2VXedicZ\nOWI7WsokgHFTsDqHO4n/3l2KTurM5p+uBZpEBaFerMgkaBVwN941I2qK/Z70lStW\nuVz0ibMGq0j2uun+zAWZjscEVkHmga05in73Gjr40XIqWflSDEI3bgze+/7QG+5s\nBDkbBrV0VHbnCjLPOgcRw2WKaARAx2Jp109OXAlXHm82a3/er6cw8kxjJKgGNg64\n/QnqY9JNXUPMzDzpUwfXHgCQliGimBsUYvWO1LH7HYSMlnhX8gxrEbUK7fErcguP\nDtlsEhuaIn/6l/IHDNSIdMGc5rS5y2jhL67Y7bg70Gy7uiEyfwXLpPOFROvy8SJa\n+FcaqI+BAgMBAAECggEACa/6WEr5P4wektCcgFXtenuwYY7BKdxKZ5cZt5CYjrbc\necLJsMUZ5guxna3A2xFIv5PJpMElWbuRrVUC8uDLHUsMafwbawbfzesGPxlppfio\nu2bF+cAC1zMModkomJzQJ1fisS5PIc4Xke6T4r8kZ/Jtsc/AHdJ6Ts7jnfpnmog9\n4zNIvPLAcw7JAoLnQePofm07KrJwwOWcRBkQliLSRd8+Wy52OvXXP8Ry2yXGW9e0\nlAgnrSAgEXc7QRbNDrGGVVsVgsMsLrBfT/8yDvDByl6laent9/PHOxYkCgWr/OtO\nYwLs7YNL3LRwzWaIzx4QLPj7ZO5LczJqDAAKGiyDsQKBgQDmAwnMCijYyRT8aysp\nuAYz1p5kchavoXKGJ5yUsFNewTR8K9ovba/mO14KkgpaqBqtvQGCKYpPxd/OjAnP\nv9Y563kWxpiI7CY9BIZUhfaqHvA6r1q9SGQBS0ZB5j72PQAeVd/TQXGGYbcNPzGf\nebx+/8QBsJlCkvA6kCEHymz60QKBgQDf385YErhi5G11WFWZrzz+gsmBCExfBIab\n0axRn+wXnfDT4ckI5DlyWiIZ02gVn200/jmN74DcpGecQUdbD6TgT8J6/NFNmC4i\n+W0085YDX3PiPzPJbyYzwzjMqTW2ItzgdeJp2UuxS/l8sq42+qozdKukgc8mkeP1\n9SRDH/wVsQKBgE7dhdNvPFgwgkCWYmNYlM/ba83XDI5F1iXHmTmmR7+6kUtuIc6X\nVnOjsXgAYQp6j0M5BjZiFemKWFXS0F5qUYLkiU1U5OI1zlqnnYOHt27XUtlcXMl/\n88I51CouTzJQ8iR0n10pGErSYFhrbZFXxVjqS4Ok0Lfx9+qslpa8QqexAoGAEcGg\njh+9/Cn9/IarE2twvQcGkHNmC0tCme1Ba5/xi9X9GfEYjtn7LHS1q7K22LAyazeW\nvQk4AUgQ57XNwQ02mIv68uJGf48IacG6xa5kQZQ6jsFQjDOCpixfvuvU1MNjHXJ8\nKMURWdiayyco5jdvdHFWg8+/7GE54XI2FBTfW6ECgYBGYiGpLD9erJbZpHsdkmHb\nIO1Hb9aYuO6f5wbe75L9ynPRbI3mdC3ISyMmvkh3rv4J2Y3GlPVrOZ65HUWMXvcS\nGsEfsyMbqO+Na/8TEML+Hz9CSzn9D/vGURYpzDuwiXhMZi8p+23FaGu3AigFQlvC\n+TTqYl4xyVlvg3WCtndvKw==\n-----END PRIVATE KEY-----\n",
  "client_email": "neodatapu2020desarrollo@neodatabases.iam.gserviceaccount.com",
  "client_id": "100961732159310987070",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/neodatapu2020desarrollo%40neodatabases.iam.gserviceaccount.com"
}`
	/*Usuarios*/
//var myDBUsuarios = "projects/neodatabases/instances/neodata/databases/neodatausuariosdesarrollo"
//var PathCredencialUsuarios = "CredentialUsuariosDesarrollo.json"


/******PRODUCCION*******/
//var myDBPU = "projects/neodatabases/instances/neodata/databases/neodatapu2020;UseClrDefaultForNull=true;Timeout=300"
//var PathCredencial = "Credential.json"

	/*Usuarios*/
var myDBUsuarios = "projects/neodatabases/instances/neodata/databases/neodatausuarios"
var PathCredencialUsuarios = `{
"type": "service_account",
"project_id": "neodatabases",
"private_key_id": "9ea1ffa856f6ceaf71f8cff436e408297ac112e5",
"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCi0belfFVSMxu0\nv+TNZKtqVGuEDI6iHwKYSe2brhoVbLitH2EKmuOuo0sjtbsgugihFk1Nw3Ym1JaQ\nBNOFsyiW3wjTqWnO7VxMFiLJ+QY4RlhPSRKzvf9MzxuSgBCGXF4swzx9xf0CZQZY\n8nC455dQ9phBKaAaCOPdcEv+M8w5pPZI/QzQAgA1Opk2IdZH6jroheDTjQM80OKy\ncIu2XwyerjbnPqoLbj6VGUWtE6Fc6QuUzp9Cx8ikKIJwMamqK48VwOIz/1a5mrwL\nMvnRsokGzKbVWpGlmJRB3RbOHnERCxwC4brTu7g4Aykmc4ZyfNycIaYe2NNJp5ns\nAMRjMTe9AgMBAAECggEABK7Ac1MhVukcvDKDWdAo6S+O3QlAasFI6raJTbHeUtFY\nnetw7sjTh7eTwmTDy19WUPtqUz13HJ1sXmRW2ydpvMA1JRoW98EWMJEA1utF76C/\nuKcq0Vme15VHHzCIxI9VAZwP+g9Td5Vp3NsWs+j1fAM4tGh9K2Fe3ATSmPe9ClnA\n/MU45SZBITJq8WaK7A4bFW7EZJqUClj1ImyL6LYXrapnI/fn7Jj2YBiaFoTVIf38\nxrLrnVuyd1AL5EPEETF5zY7mlgEXAHNEGBz/1Qvzdc4CVxHdzD4VWsskV2+4um+e\nko2LaXU3Y7ogZCKq2lI5ZyKh6dWUahxVK+b8yDXVqQKBgQDPLIbU1UnxiP/y6hPd\nMYClRUp5HSQ7smHkSzPsrTx16BDinJc6Mdk3XszT+C3kxOXvPAZoo6RRXfM6iuHE\nu7aSvEc++3IHsyKJG5DrIcBob+dw6/IhhQD+qfJeaW7CDDtSt8oe0IyKh0xoZnDM\nDx1RGSR2AzPfA9kVECsJK9vPlQKBgQDJMSBy7JF+l32MEIIibpnITbzlay11hO7c\n4YyGzhJgulawfQ6btJONZ+fhwgFcVtwKQJk3F9eDEey5ZU7DOpvA8nNp8jJWGNx8\n2WnYHpPyzVg94+iBxPUkidXkX9invtcmDQK1E3rkCzA/u3AokMaZNv3rKxhTScZo\nsTkzXg5diQKBgQCV3W/rph7HSoNcBVg5OgYFiApHcn0zcea+bmgxXV/PKwCL7Rq4\nTjcDhuBEtdL9GjH6zUEVkMrc0UmRn8WRuvhnTSt8tAfY6gyhgf9jGBdHdSvG0+dj\nuALwNZ/ODBbC6HZBcBIWGGaSDy2GmBQgiQ58C+QEOmzbdr0iL3/yMnffwQKBgQCu\n+mEEUqc/eDWiqYDkvVhWEvYkeZBx0wmDZU64t2TYZ3eZy0n3NZfWtfXALODOFGUP\nLZuThNLUlbRSkb9sn/5yUur5y8DnjHvGwbgCVKXL17fVK/A9XLTv8EjsdEeTrLCl\n0U73eVe6Gdj+tOAZB8ER4/f2neZsGY/L4caj3DuWMQKBgHW+0pE4FiHtnNG6EdCB\nQl3D7WrBDBbIjWKtlM6p+C443OB39lTuy+ZIxaGOPr+nz02VgtPqWqgs046F5gjS\npBuLSZTnnPAQqSOiWLskcWqV6NIGzHD38y5iBdlF2Lp/mKsutnTwbyjQGDFABFyD\nzbN3Tsbh0dFa5/lzPzD9+rtL\n-----END PRIVATE KEY-----\n",
"client_email": "neodatausuarios3@neodatabases.iam.gserviceaccount.com",
"client_id": "105646190877296060833",
"auth_uri": "https://accounts.google.com/o/oauth2/auth",
"token_uri": "https://oauth2.googleapis.com/token",
"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/neodatausuarios3%40neodatabases.iam.gserviceaccount.com"
}`



func ConexionPU() (client *spanner.Client, ctx context.Context, error clases.ClsErrores){

	ctx = context.Background()
	client, err := spanner.NewClient(ctx, myDBPU, option.WithCredentialsJSON([]byte(PathCredencial)))
	if err != nil {
		error.Error = "ConexionError"
		error.ErrorDescripcion = "Error al abrir la base de datos " + err.Error()
	}
	return client, ctx, error
}

func ConexionUsuarios() (client *spanner.Client, ctx context.Context, error clases.ClsErrores){

	ctx = context.Background()
	client, err := spanner.NewClient(ctx, myDBUsuarios, option.WithCredentialsFile(PathCredencialUsuarios))
	if err != nil {
		error.Error = "ConexionError"
		error.ErrorDescripcion = "Error al abrir la base de datos " + err.Error()
	}
	return client, ctx, error
}
