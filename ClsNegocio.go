package negocio

import (
	"cloud.google.com/go/spanner"
	"context"
	"github.com/jaimemr86/clases"
	"google.golang.org/api/iterator"
)

var client *spanner.Client

func CatalogoRegistraActualiza(ListaCatalogo []clases.ClsCatalogo, ctx context.Context) (result clases.ClsRegresaCatalogo) {

	lsDic := make(map[int64]int64)
	var idCodigo int64

	for _, obj := range ListaCatalogo {
		idCodigo = 0
		if obj.IdCodigoNube > 0 {
			if !obj.NoActualizaCatalogo {
				m := spanner.Update("Catalogo",
					[]string{"IdCatalogoDeObras", "IdCodigo", "Codigo", "CodigoSap", "Descripcion", "DescripcionLarga",
						"EsAgrupador", "EsPorcentaje", "IdFamilia", "IdFichaTecnica", "IdImagen", "IdProcedimiento", "IdProveedor",
						"IdTipoInsumo", "IdUnidad", "InsumoDescontinuado", "PorcentajeFondoGarantia", "Referencia", "VolumenDefault"},
					[]interface{}{obj.IdCatalogoDeObras, obj.IdCodigoNube, obj.Codigo, obj.CodigoSap, obj.Descripcion, obj.DescripcionLarga, obj.EsAgrupador, obj.EsPorcentaje,
						obj.IdFamilia, obj.IdFichaTecnica, obj.IdImagen, obj.IdProcedimiento, obj.IdProveedor, obj.IdTipoInsumo, obj.IdUnidad, obj.InsumoDescontinuado,
						obj.PorcentajeFondoGarantia, obj.Referencia, obj.VolumenDefault})
				_, err := client.Apply(ctx, []*spanner.Mutation{m})
				if err != nil {
					result.Errores.Error = "UpdateError"
					result.Errores.ErrorDescripcion = err.Error()
					goto ResErrores
				}
			}

		} else {

			stmt := spanner.NewStatement(`SELECT IdCodigo FROM Catalogo WHERE Codigo = @Codigo AND IdCatalogoDeObras = @IdCatalogoDeObras`)
			stmt.Params["Codigo"] = obj.Codigo
			stmt.Params["IdCatalogoDeObras"] = obj.IdCatalogoDeObras

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
				if err := row.Columns(&idCodigo); err != nil {
					result.Errores.Error = "SelectError2"
					result.Errores.ErrorDescripcion = err.Error()
					goto ResErrores
				}
			}

			if idCodigo == 0 {
				idCodigo = int64(GeneraGuid(0))
				m := spanner.Insert("Catalogo",
					[]string{"IdCatalogoDeObras", "IdCodigo", "Codigo", "CodigoSap", "Descripcion", "DescripcionLarga",
						"EsAgrupador", "EsPorcentaje", "IdFamilia", "IdFichaTecnica", "IdImagen", "IdProcedimiento", "IdProveedor",
						"IdTipoInsumo", "IdUnidad", "InsumoDescontinuado", "PorcentajeFondoGarantia", "Referencia", "VolumenDefault"},
					[]interface{}{obj.IdCatalogoDeObras, idCodigo, obj.Codigo, obj.CodigoSap, obj.Descripcion, obj.DescripcionLarga, obj.EsAgrupador, obj.EsPorcentaje,
						obj.IdFamilia, obj.IdFichaTecnica, obj.IdImagen, obj.IdProcedimiento, obj.IdProveedor, obj.IdTipoInsumo, obj.IdUnidad, obj.InsumoDescontinuado,
						obj.PorcentajeFondoGarantia, obj.Referencia, obj.VolumenDefault})

				_, err := client.Apply(ctx, []*spanner.Mutation{m})
				if err != nil {
					result.Errores.Error = "InsertError"
					result.Errores.ErrorDescripcion = err.Error()
					goto ResErrores
				}
			} else {
				if !obj.NoActualizaCatalogo {
					m := spanner.Update("Catalogo",
						[]string{"IdCatalogoDeObras", "IdCodigo", "Codigo", "CodigoSap", "Descripcion", "DescripcionLarga",
							"EsAgrupador", "EsPorcentaje", "IdFamilia", "IdFichaTecnica", "IdImagen", "IdProcedimiento", "IdProveedor",
							"IdTipoInsumo", "IdUnidad", "InsumoDescontinuado", "PorcentajeFondoGarantia", "Referencia", "VolumenDefault"},
						[]interface{}{obj.IdCatalogoDeObras, idCodigo, obj.Codigo, obj.CodigoSap, obj.Descripcion, obj.DescripcionLarga, obj.EsAgrupador, obj.EsPorcentaje,
							obj.IdFamilia, obj.IdFichaTecnica, obj.IdImagen, obj.IdProcedimiento, obj.IdProveedor, obj.IdTipoInsumo, obj.IdUnidad, obj.InsumoDescontinuado,
							obj.PorcentajeFondoGarantia, obj.Referencia, obj.VolumenDefault})
					_, err := client.Apply(ctx, []*spanner.Mutation{m})
					if err != nil {
						result.Errores.Error = "UpdateError"
						result.Errores.ErrorDescripcion = err.Error()
						goto ResErrores
					}
				}
			}
			lsDic[obj.IdCodigoSql] = idCodigo
			result.ListaIds = lsDic
		}
	}
ResErrores:

	return result
}
