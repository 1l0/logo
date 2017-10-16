package main

import (
	"os/exec"

	"github.com/fogleman/gg"
)

const (
	// Font paths
	fontRegular = `/Library/Fonts/NotoSans-Regular.ttf`
	fontBold    = `/Library/Fonts/NotoSans-Bold.ttf`
	// Size
	S = 640
)

var (
	// Context
	C = gg.NewContext(S, S)
	// Params
	P = map[string][]float64{
		"logo":       {S / 32, -S * .13},
		"text":       {S * .5, S * .77},
		"fontsize":   {S * .155},
		"textAnc":    {-S * .076, 0},
		"dot":        {S / 16},
		"line":       {S * .025},
		"root":       {S*.5 + S*.343, S * .5},
		"edgeTop":    {S * .53, S * .347},
		"edgeBottom": {S * .53, S * .632},
		"midTop":     {S * .689, S * .4234},
		"midBottom":  {S * .689, S * .565},
		"edgeMid":    {S * .53, S * .49},
		"eccTop":     {S * .689, S * .245},
		"eccBottom":  {S * .689, S * .733},
	}
)

func main() {
	C.SetRGBA(0, 0, 0, 1)
	DrawLessthan(false)
	//DrawLessthan(true)
	DrawDot(false)
	DrawLine(false)
	drawText(false)
	test()
}

func DrawLine(debug bool) {
	C.Push()
	defer C.Pop()
	C.Translate(P["logo"][0], P["logo"][1])
	C.SetLineWidth(P["line"][0])
	C.MoveTo(P["root"][0], P["root"][1])
	C.LineTo(P["edgeTop"][0], P["edgeTop"][1])
	C.MoveTo(P["root"][0], P["root"][1])
	C.LineTo(P["edgeBottom"][0], P["edgeBottom"][1])
	C.MoveTo(P["midBottom"][0], P["midBottom"][1])
	C.LineTo(P["edgeMid"][0], P["edgeMid"][1])
	/*C.MoveTo(P["midTop"][0], P["midTop"][1])
	C.LineTo(P["eccTop"][0], P["eccTop"][1])*/
	C.MoveTo(P["midBottom"][0], P["midBottom"][1])
	C.LineTo(P["eccBottom"][0], P["eccBottom"][1])
	C.Stroke()
}

func DrawDot(debug bool) {
	C.Push()
	defer C.Pop()
	if debug {
		C.SetRGBA(0, 0, 0, .5)
	}
	C.Translate(P["logo"][0], P["logo"][1])
	C.DrawCircle(P["root"][0], P["root"][1], P["dot"][0])
	C.DrawCircle(P["edgeTop"][0], P["edgeTop"][1], P["dot"][0])
	C.DrawCircle(P["edgeBottom"][0], P["edgeBottom"][1], P["dot"][0])
	C.DrawCircle(P["midTop"][0], P["midTop"][1], P["dot"][0])
	C.DrawCircle(P["midBottom"][0], P["midBottom"][1], P["dot"][0])
	C.DrawCircle(P["edgeMid"][0], P["edgeMid"][1], P["dot"][0])
	//C.DrawCircle(P["eccTop"][0], P["eccTop"][1], P["dot"][0])
	C.DrawCircle(P["eccBottom"][0], P["eccBottom"][1], P["dot"][0])
	C.Fill()
}

func DrawLessthan(debug bool) {
	C.Push()
	defer C.Pop()
	if debug {
		C.Translate(P["logo"][0], P["logo"][1]-S*0.0182)
		C.SetRGBA(0, 0, 0, 0.1)
	} else {
		C.Translate(-P["logo"][0], P["logo"][1]-S*0.0182)
	}
	if err := C.LoadFontFace(fontBold, S*.7); err != nil {
		panic(err)
	}
	if debug {
		C.DrawStringAnchored(">", S/2, S/2, 0, .5)
	} else {
		C.DrawStringAnchored("<", S/2, S/2, 1, .5)
	}
}

func drawText(debug bool) {
	C.Push()
	defer C.Pop()
	C.Translate(P["text"][0], P["text"][1])
	if debug {
		C.SetRGBA(0, 0, 0, 0.3)
	} else {
		C.SetRGBA(0, 0, 0, 1)
	}
	/*if err := C.LoadFontFace(fontBold, P["fontsize"][0]); err != nil {
		panic(err)
	}
	C.DrawStringAnchored("W", P["textAnc"][0], P["textAnc"][1], 1, .5)
	if err := C.LoadFontFace(fontRegular, P["fontsize"][0]); err != nil {
		panic(err)
	}
	C.DrawStringAnchored("ASM", P["textAnc"][0], P["textAnc"][1], 0, .5)*/
	if err := C.LoadFontFace(fontBold, P["fontsize"][0]); err != nil {
		panic(err)
	}
	C.DrawStringAnchored("WASM", 0, 0, .5, .5)
}

func test() {
	file := "wasm.png"
	if err := C.SavePNG(file); err != nil {
		panic(err)
	}
	if err := exec.Command("open", file).Run(); err != nil {
		panic(err)
	}
}
