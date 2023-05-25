package dpfm_api_output_formatter

import (
	"data-platform-api-product-master-doc-creates-rmq-kube/DPFM_API_Caller/requests"
)

func ConvertToGeneralDoc(generalDoc *GeneralDoc) *requests.GeneralDoc {
	pm := &requests.GeneralDoc{}

	pm = &requests.GeneralDoc{
		Product:                  generalDoc.Product,
		DocType:                  generalDoc.DocType,
		DocVersionID:             generalDoc.DocVersionID,
		DocID:                    generalDoc.DocID,
		FileExtension:            generalDoc.FileExtension,
		FileName:                 generalDoc.FileName,
		FilePath:                 generalDoc.FilePath,
		DocIssuerBusinessPartner: generalDoc.DocIssuerBusinessPartner,
	}

	return pm
}
