package negocio

import (
	"github.com/jaimemr86/clases"
	"cloud.google.com/go/spanner"
	"context"
	"google.golang.org/api/option"
)

/******DESARROLLO*******/
var myDBPU = "projects/neodatabases/instances/neodata/databases/neodatapu2020desarrollo"
var PathCredencial = "negocio/CredentialDesarrollo.json"

	/*Usuarios*/
//var myDBUsuarios = "projects/neodatabases/instances/neodata/databases/neodatausuariosdesarrollo"
//var PathCredencialUsuarios = "CredentialUsuariosDesarrollo.json"


/******PRODUCCION*******/
//var myDBPU = "projects/neodatabases/instances/neodata/databases/neodatapu2020;UseClrDefaultForNull=true;Timeout=300"
//var PathCredencial = "Credential.json"

	/*Usuarios*/
var myDBUsuarios = "projects/neodatabases/instances/neodata/databases/neodatausuarios"
var PathCredencialUsuarios = "CredentialUsuarios.json"



func ConexionPU() (client *spanner.Client, ctx context.Context, error clases.ClsErrores){

	ctx = context.Background()
	client, err := spanner.NewClient(ctx, myDBPU, option.WithCredentialsFile(PathCredencial))
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
