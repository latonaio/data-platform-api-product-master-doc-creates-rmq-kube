package dpfm_api_input_reader

import (
	"data-platform-api-product-master-doc-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneralDoc() *requests.GeneralDoc {
	data := sdc.Product.GeneralDoc
	return &requests.GeneralDoc{
		Product:                  data.Product,
		DocType:                  data.DocType,
		DocVersionID:             data.DocVersionID,
		DocID:                    data.DocID,
		FileExtension:            data.FileExtension,
		FileName:                 data.FileName,
		FilePath:                 data.FilePath,
		DocIssuerBusinessPartner: data.DocIssuerBusinessPartner,
	}
}
