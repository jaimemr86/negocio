package negocio

import (
	"bytes"
	"cloud.google.com/go/civil"
	"cloud.google.com/go/spanner"
	"context"
	"encoding/json"
	"github.com/jaimemr86/clases"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"time"
)

const PROVEEDOR_GOOGLE = "google.com"
const GOOGLE_REDIRECT_URL = "https://neodata-usuarios-245016.web.app/__/auth/handler"
const URLPOSTFIREBASE = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyAssertion"
const APIKEY = "AIzaSyAS2to3Z1LwDb-RatRgXth3thXYRLtkG6I"

func ObtieneAccessToken(accessToken string, client *spanner.Client, ctx context.Context) (IdUsuario int64) {

	FechaAccessToken := time.Now().UTC()
	//FechaActual := time.Now().UTC()

	if IdUsuario == 0 {
		stmt := spanner.NewStatement(`SELECT IdUsuario,FechaHora FROM AccessTokens WHERE AccessToken = @AccessToken `)
		stmt.Params["AccessToken"] = accessToken

		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				IdUsuario = 0
				goto ResErrores
			}
			if err := row.Columns(&IdUsuario, &FechaAccessToken); err != nil {
				IdUsuario = 0
				goto ResErrores
			}
		}
		/*after := FechaAccessToken.Add(50 * time.Minute)
		if after.Unix() < FechaActual.Unix() {
			IdUsuario = 0
		}*/
	}

ResErrores:

	return IdUsuario
}

func ObtieneDatosUsuarioDatosToken(accessToken string) (result clases.FirebaseUserPerfil) {

	postBodyTmp := "access_token=" + accessToken + "&providerId=" + PROVEEDOR_GOOGLE

	objFirebase := clases.FirebaseVerifyAssertion{
		PostBody:            postBodyTmp,
		RequestUri:          GOOGLE_REDIRECT_URL,
		ReturnIdpCredential: true,
		ReturnSecureToken:   true,
	}
	jsonValue, _ := json.Marshal(objFirebase)
	response, err := http.Post(URLPOSTFIREBASE+"?key="+APIKEY, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		result.Errores.Error = "Error ObtieneDatosUsuarioDatosToken"
		result.Errores.ErrorDescripcion = err.Error()
		goto ResErrores
	} else {
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			result.Errores.Error = "Error ObtieneDatosUsuarioDatosToken"
			result.Errores.ErrorDescripcion = err.Error()
			goto ResErrores
		}
		if response.StatusCode != 200 {
			result.Errores.Error = response.Status
			result.Errores.ErrorDescripcion = "Token no válido"
			goto ResErrores
		}
		errJ := json.Unmarshal(responseData, &result)
		if errJ != nil {
			result.Errores.Error = "Error ObtieneDatosUsuarioDatosToken"
			result.Errores.ErrorDescripcion = errJ.Error()
			goto ResErrores
		}
	}
ResErrores:

	return result
}

func ConfirmaSesionUsuarioAdministrador(email string, codigoSistema string, idSesion int64, IdUsuario int64, client *spanner.Client, ctx context.Context) (result clases.ClsDatosCliente) {

	var objUsu clases.ClsDatosCliente
	var datosLicenciaCliente clases.ClsDatosLicenciaCliente
	var objSesion clases.ClsSesion
	var fechaActual civil.Date

	datosLicenciaCliente = ObtieneLicenciaCliente(email, codigoSistema, IdUsuario, client, ctx)

	objSesion = ConfirmaSesionActiva(idSesion, client, ctx)

	objUsu.AccessToken = "" //objRefreshToken.access_token;
	objUsu.Sesion = objSesion.IdSesionActiva
	objUsu.NumeroLicencia = datosLicenciaCliente.NumeroLicencia
	objUsu.RazonSocial = datosLicenciaCliente.RazonSocial
	objUsu.FechaVigencia = datosLicenciaCliente.FechaVigencia.String()
	objUsu.RefreshToken = "" //refeshToken;
	objUsu.OtraSesionActiva = objSesion.TieneActiva
	objUsu.NoTieneVigencia = false
	objUsu.CambioIp = false
	objUsu.CaducoSesion = false
	objUsu.TokenCaducado = false
	objUsu.IdUsuario = datosLicenciaCliente.IdUsuario
	objUsu.LicenciaEstudiantil = datosLicenciaCliente.LicenciaEstudiantil

	if objSesion.IdSesionActiva > 0 {
		if fechaActual.String() > datosLicenciaCliente.FechaVigencia.String() {
			objUsu.RazonSocial = "Versión de demostración"
			objUsu.FechaVigencia = "Sin vigencia"
			objUsu.NoTieneVigencia = true
		}
		if datosLicenciaCliente.NumeroLicencia <= 0 {
			objUsu.RazonSocial = "Versión de demostración"
			objUsu.FechaVigencia = "Sin vigencia"
			objUsu.NoTieneVigencia = true
		}
	} else {
		objUsu.RazonSocial = "Versión de demostración"
		objUsu.CaducoSesion = true
		objUsu.Errores.Error = "Error ConfirmaSesionUsuarioAdministrador"
		objUsu.Errores.ErrorDescripcion = "La sesion a caducado"
	}
	return objUsu
}

func ObtieneLicenciaCliente(email string, codigoSistema string, idUsuario int64, client *spanner.Client, ctx context.Context) (result clases.ClsDatosLicenciaCliente) {

	if idUsuario == 0 {
		stmt := spanner.NewStatement(`SELECT IdUsuario, RefreshToken FROM Usuarios WHERE Usuarios.Emailusuario = @email `)
		stmt.Params["email"] = email

		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				result.Errores.Error = "SelectError1"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			if err := row.Columns(&result.IdUsuario, &result.RefreshToken); err != nil {
				result.Errores.Error = "SelectError2"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			idUsuario = result.IdUsuario
		}

	} else {
		result.IdUsuario = idUsuario
	}

	if result.Errores.Error == "" {
		stmt := spanner.NewStatement("SELECT Licencias.IdLicencia IdLicencia, Sistemas.IdSistema IdSistema, Licencias.IdUsuario IdUsuario, " +
			"Licencias.NumeroLicencia NumeroLicencia, Licencias.FechaVigencia FechaVigencia, " +
			"Empresas.RazonSocial RazonSocial, Usuarios.RefreshToken RefreshToken, " +
			"IFNULL(Empresas.LicenciaEstudiantil,false) as LicenciaEstudiantil " +
			"FROM Usuarios " +
			"LEFT join@{JOIN_TYPE=APPLY_JOIN} Licencias on Usuarios.IdUsuario = Licencias.IdUsuario " +
			"LEFT join@{JOIN_TYPE=APPLY_JOIN} Sistemas on Licencias.idSistema = Sistemas.idSistema " +
			"LEFT Join@{JOIN_TYPE=APPLY_JOIN} Empresas on Licencias.IdEmpresa = Empresas.IdEmpresa " +
			"WHERE Usuarios.IdUsuario = @IdUsuario And Sistemas.CodigoSistema = @CodigoSistema ")
		stmt.Params["IdUsuario"] = idUsuario
		stmt.Params["CodigoSistema"] = codigoSistema

		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				result.Errores.Error = "SelectError1"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			if err := row.Columns(&result.IdLicencia, &result.IdSistema,
				&result.IdUsuario, &result.NumeroLicencia,
				&result.FechaVigencia, &result.RazonSocial,
				&result.RefreshToken, &result.LicenciaEstudiantil); err != nil {
				result.Errores.Error = "SelectError2"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
		}
	}
ResErrores:

	return result
}

func ConfirmaSesionActiva(idSesion int64, client *spanner.Client, ctx context.Context) (result clases.ClsSesion) {

	result.TieneActiva = true

	ActualizaUltimaLlamada(idSesion, client, ctx)

	if result.Errores.Error == "" {
		stmt := spanner.NewStatement(`SELECT IdSesion, IpPublica FROM Sesiones WHERE IdSesion = @idSesion AND Activa = true`)
		stmt.Params["IdSesion"] = idSesion
		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				result.Errores.Error = "SelectError1"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			if err := row.Columns(&result.IdSesionActiva, &result.IpPublica); err != nil {
				result.Errores.Error = "SelectError2"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			result.TieneActiva = false
		}
	}
ResErrores:

	return result
}

func ActualizaUltimaLlamada(idSesion int64, client *spanner.Client, ctx context.Context) (result clases.ClsSesion) {

	var exampleTimestamp = time.Now()

	if result.Errores.Error == "" {
		stmt := spanner.Statement{SQL: `UPDATE Sesiones SET FechaUltimaLlamada = @FechaUltimaLlamada WHERE IdSesion = @IdSesion`,
			Params: map[string]interface{}{
				"FechaUltimaLlamada": exampleTimestamp,
				"IdSesion":           idSesion,
			},
		}
		_, err := client.PartitionedUpdate(ctx, stmt)
		if err != nil {
			result.Errores.Error = "UpdateError"
			result.Errores.ErrorDescripcion = err.Error()
			goto ResErrores
		}
		result.IdSesionActiva = idSesion
	}

ResErrores:

	return result
}

func ObtieneUsuarioAdministrador(IdUsuario int64, codigoSistema string, client *spanner.Client, ctx context.Context) (result clases.ClsUsuarioAdmin) {

	if result.Errores.Error == "" {
		stmt := spanner.NewStatement("SELECT Usuarios.IdUsuario,Usuarios.EmailUsuario,Empresas.IdUsuarioAdministrador AS IdUsuarioAdmin," +
			"UsuarioAdmin.EmailUsuario AS EmailUsuarioAdmin,Empresas.IdEmpresa," +
			"IFNULL(Empresas.LicenciaEstudiantil,false) as LicenciaEstudiantil " +
			"FROM Usuarios " +
			"HASH JOIN Licencias ON Usuarios.IdUsuario = Licencias.IdUsuario " +
			"HASH JOIN Empresas ON Licencias.IdEmpresa = Empresas.IdEmpresa " +
			"HASH JOIN Sistemas ON Licencias.IdSistema = sistemas.IdSistema " +
			"HASH JOIN Usuarios AS UsuarioAdmin ON Empresas.IdUsuarioAdministrador = UsuarioAdmin.IdUsuario " +
			"WHERE Usuarios.IdUsuario = @IdUsuario AND Sistemas.CodigoSistema = @codigoSistema")
		stmt.Params["IdUsuario"] = IdUsuario
		stmt.Params["codigoSistema"] = codigoSistema

		iter := client.Single().Query(ctx, stmt)
		defer iter.Stop()
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				result.Errores.Error = "SelectError1"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
			if err := row.Columns(
				&result.IdUsuario,
				&result.Email,
				&result.IdUsuarioAdmin,
				&result.EmailAdmin,
				&result.IdEmpresa,
				&result.LicenciaEstudiantil); err != nil {
				result.Errores.Error = "SelectError2"
				result.Errores.ErrorDescripcion = err.Error()
				goto ResErrores
			}
		}
	}
ResErrores:

	return result
}

func ObtieneUsuarioAdmin(accessToken clases.ClsAccessToken, client *spanner.Client, ctx context.Context) (result clases.ClsDatosCliente) {

	IdUsuario := ObtieneAccessToken(accessToken.AccessToken, client, ctx)

	if IdUsuario == 0 {
		objF := ObtieneDatosUsuarioDatosToken(accessToken.AccessToken)
		if objF.Errores.Error == "" {
			result = ConfirmaSesionUsuarioAdministrador(objF.Email, accessToken.CodigoSistema, accessToken.IdSesion, IdUsuario, client, ctx)
			objUsuarioAdmin := ObtieneUsuarioAdministrador(result.IdUsuario, accessToken.CodigoSistema, client, ctx)
			result.IdUsuario = objUsuarioAdmin.IdUsuario
			result.IdUsuarioAdmin = objUsuarioAdmin.IdUsuarioAdmin
			result.EmailAdmin = objUsuarioAdmin.EmailAdmin
			result.IdEmpresa = objUsuarioAdmin.IdEmpresa
			result.Email = objUsuarioAdmin.Email
			result.LicenciaEstudiantil = objUsuarioAdmin.LicenciaEstudiantil
		} else {
			result.RazonSocial = "Versión de demostración"
			result.TokenCaducado = true
			result.Errores.Error = "Error ObtieneUsuarioAdmin"
			result.Errores.ErrorDescripcion = "Token caducado"
		}
	} else {
		result = ConfirmaSesionUsuarioAdministrador("", accessToken.CodigoSistema, accessToken.IdSesion, IdUsuario, client, ctx)
		objUsuarioAdmin := ObtieneUsuarioAdministrador(IdUsuario, accessToken.CodigoSistema, client, ctx)
		result.IdUsuario = objUsuarioAdmin.IdUsuario
		result.IdUsuarioAdmin = objUsuarioAdmin.IdUsuarioAdmin
		result.EmailAdmin = objUsuarioAdmin.EmailAdmin
		result.IdEmpresa = objUsuarioAdmin.IdEmpresa
		result.Email = objUsuarioAdmin.Email
		result.LicenciaEstudiantil = objUsuarioAdmin.LicenciaEstudiantil
	}
	return result
}

func ObtieneUsuario(objCode clases.ClsAccessToken, SePermiteDemo bool, client *spanner.Client, ctx context.Context) (result clases.ClsDatosCliente) {

	result = ObtieneUsuarioAdmin(objCode, client, ctx)
	var usuarioOk bool

	if len(result.Errores.Error) > 0 {
		usuarioOk = false
	}
	if usuarioOk && result.CaducoSesion {
		result.Errores.Error = "Error"
		result.Errores.ErrorDescripcion = "Caducó la sesión"
		usuarioOk = false
	}
	if !SePermiteDemo && usuarioOk && result.NumeroLicencia == 0 {
		result.Errores.Error = "Error"
		result.Errores.ErrorDescripcion = "El usuario es DEMO."
		usuarioOk = false
	}
	if !SePermiteDemo && usuarioOk && result.NoTieneVigencia {
		result.Errores.Error = "Error"
		result.Errores.ErrorDescripcion = "La licencia del usuario ya expiró."
		usuarioOk = false
	}
	if usuarioOk && result.TokenCaducado {
		result.Errores.Error = "Error"
		result.Errores.ErrorDescripcion = "El token del usuario ya caducó."
		usuarioOk = false
	}
	if usuarioOk && result.OtraSesionActiva {
		result.Errores.Error = "Error"
		result.Errores.ErrorDescripcion = "El usuario tiene otra sesión activa."
		usuarioOk = false
	}

	return result
}
