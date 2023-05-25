package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-master-doc-creates-rmq-kube/DPFM_API_Input_Formatter"
	dpfm_api_output_formatter "data-platform-api-product-master-doc-creates-rmq-kube/DPFM_API_Output_Formatter"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	generalDoc *dpfm_api_output_formatter.GeneralDoc,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	c.generalDocCreateSql(nil, input, output, generalDoc, errs, log)
	response := dpfm_api_output_formatter.ConvertToGeneralDoc(generalDoc)

	data := dpfm_api_output_formatter.Message{
		GeneralDoc: dpfm_api_output_formatter.GeneralDoc{
			Product:                  response.Product,
			DocType:                  response.DocType,
			DocVersionID:             response.DocVersionID,
			DocID:                    response.DocID,
			FileExtension:            response.FileExtension,
			FileName:                 response.FileName,
			FilePath:                 response.FilePath,
			DocIssuerBusinessPartner: response.DocIssuerBusinessPartner,
		},
	}

	return data
}

func (c *DPFMAPICaller) generalDocCreateSql(
	ctx context.Context,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	generalDoc *dpfm_api_output_formatter.GeneralDoc,
	errs *[]error,
	log *logger.Logger,
) {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalDoc, "function": "ProductMaterDoc", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		*errs = append(*errs, err)
		return
	}
	res.Success()

	return
}
