package render_test

import (
	"github.com/payboxth/goslip"
	"testing"
	"time"

	"github.com/payboxth/goslip/render"
	"github.com/stretchr/testify/assert"
)

func TestNewRender(t *testing.T) {
	r := render.New()
	assert.NotNil(t, r, "NewPrinter should not nil: %v", r)
}

// Test HTMLtoSlipJPG must install wkhtml to host server by
func TestRender_Bytes1(t *testing.T) {
	r := render.New()
	b, err := r.Bytes(html, 561, "jpg")
	if err != nil {
		t.Errorf("printer.Bytes() error: %v", err)
	}
	assert.NotNil(t, b, "Bytes() should not return nil")
}

func TestRender_Bytes2(t *testing.T) {
	r := render.New()
	b, err := r.Bytes2(body, html, 561, "jpg")
	if err != nil {
		t.Errorf("Error cal render.Byte2(): %v", err)
	}
	assert.NotNil(t, b, "Bytes2() should not return nil")
}

var html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>eSlip</title>
    <meta name="viewport" content="width=device-width"/>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>
<body>
    <div id="capture" class="body" style="padding: 10px; background: #ffffff;width:540px;height:auto;">
        <h4 style="color: #000;text-align: center;color:#000">TAX INV (ABB)</h4>
        <div class="row">

            <div class="text-left" style="float: left;">NO: {{.DocNo}}</div>
            <div style="text-align:right">DATE: {{.Date}}</div>
        </div>
        <hr>
        <h4 style="color: #000;text-align: center;    font-size: 1.7em;">บริษัท นพดลพานิช จำกัด</h4>
        <hr>
        <div class="row">
            <div class="text-left" style="float: left;">TAX ID: {{.TaxID}}</div>
            <div style="text-align:right">POS: {{.Pos}}</div>
            <div class="text-left" style="float: left;">CS: {{.CS}}</div>
            <div style="text-align:right">TIME: {{.Time}}</div>
            <div class="row text-left" style="float: left;">พนักงาน: PAYBOX</div>
            <div style="clear: both;"></div>
        </div>
        <hr>
        <div class="row" style="font-size:1.6em;    font-weight: 900;">
            <div class="text-left" style="float: left;"> รายการ</div>
            <div style="text-align:right"> มูลค่า</div>
        </div>
        <div style="clear: both;"></div>
        <div class="row" style="font-size:1.3em;margin-top:5px;margin-bottom:5px;">
            {{range $index, $element := .SaleSubs}}

            <div class="text-left" style="float: left;width: 72%;word-wrap:break-word;">
                <div style="float:left">{{$element.ItemName}} x {{$element.Qty}}</div>
                <div style="float:left;padding-left:5px;"></div>
            </div>
            <div style="float: right;text-align:right;width: 27%;word-wrap:break-word;">{{$element.Price}}</div>
            <div style="margin-top:10px;clear: both;"></div>

            {{end}}
            <br>
        </div>


        <hr>
        <div class="row" style="font-size:1.2em;">
            <div class="text-left" style="float: left;">TOTAL:</div>
            <div style="text-align:right">{{.Total}}</div>
            <div class="text-left" style="float: left;">VAT AMOUNT:7%</div>
            <div style="text-align:right">{{.Vat}}</div>
            <div class="text-left" style="float: left;">CREDIT CARD:</div>
            <div style="text-align:right">{{.CreditCard}}</div>
            <div class="text-left" style="float: left;">VAT INCLUDED:</div>
            <div style="text-align:right">CHANGE: 0</div>

            <div class="col-12" style="text-align: center;">ขอบคุณที่มาใช้บริการ</div>
            <br>
            <div class="col-12" style="text-align: center;">สอบถามโทร. {{.Phone}} 8.00-17.00 ทุกวัน</div>

        </div>
    </div>

</body>
<style>
    @import 'https://fonts.googleapis.com/css?family=Kanit|Prompt|Athiti|Pridi';
    .body{
        font-family: Kanit;
        font-style: normal;
        font-size: 1em;
    }
    .row {
        width: 100%;
        font-family: Athiti;
        font-style: normal;
        font-size: 1em;
    }

    .col, .col-1, .col-10, .col-11, .col-12, .col-2, .col-3, .col-4, .col-5, .col-6, .col-7, .col-8, .col-9, .col-auto, .col-lg, .col-lg-1, .col-lg-10, .col-lg-11, .col-lg-12, .col-lg-2, .col-lg-3, .col-lg-4, .col-lg-5, .col-lg-6, .col-lg-7, .col-lg-8, .col-lg-9, .col-lg-auto, .col-md, .col-md-1, .col-md-10, .col-md-11, .col-md-12, .col-md-2, .col-md-3, .col-md-4, .col-md-5, .col-md-6, .col-md-7, .col-md-8, .col-md-9, .col-md-auto, .col-sm, .col-sm-1, .col-sm-10, .col-sm-11, .col-sm-12, .col-sm-2, .col-sm-3, .col-sm-4, .col-sm-5, .col-sm-6, .col-sm-7, .col-sm-8, .col-sm-9, .col-sm-auto, .col-xl, .col-xl-1, .col-xl-10, .col-xl-11, .col-xl-12, .col-xl-2, .col-xl-3, .col-xl-4, .col-xl-5, .col-xl-6, .col-xl-7, .col-xl-8, .col-xl-9, .col-xl-auto {
        position: relative;
        width: 100%;
        padding-right: 5px;
        padding-left: 5px;
    }

    .cols-10 {
        width: 350px;
    }

    .cols-2 {
        width: 90px;
    }

    .col-6 {
        -ms-flex: 0 0 50%;
        flex: 0 0 50%;
        max-width: calc(50% - 10px);
    }

    .col-12 {
        -ms-flex: 0 0 100%;
        flex: 0 0 100%;
        max-width: 100%;
    }
</style>
</html>
`

var body = &slip.Body{
	Title:      "MakeKAFE",
	DocDate:    "2021-04-02T15:04:05+07:00", // time formated in RFC3339
	DocNumber:  "000001",
	Ref:        "202341",
	CreateDate: time.Now().Format(time.RFC3339),
	Lines: []slip.Line{
		{
			Seq:   1,
			SKU:   "123456",
			Name:  "Product Name1",
			Qty:   1.0,
			Price: 100,
			Note:  "",
		},
		{
			Seq:   2,
			SKU:   "78910",
			Name:  "Product Name2",
			Qty:   2.0,
			Price: 50,
			Note:  "Test Note",
		},
	},
}
