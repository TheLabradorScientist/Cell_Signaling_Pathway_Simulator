package main

import (
	"image/color"
	"math/rand"
	"github.com/hajimehoshi/ebiten/v2"
)

type ReceptionLevel struct {
	// PLASMA SPRITES
	protoPlasmaBg  		StillImage
	plasmaBg       		Parallax
	plasmaMembrane 		Parallax
	signal         		Signal
	receptorA      		Receptor
	receptorB      		Receptor
	receptorC     		Receptor
	receptorD      		Receptor
	temp_tk1A      		Kinase
	temp_tk1B			Kinase
	temp_tk1C      		Kinase
	temp_tk1D     		Kinase
	infoButton			InfoPage
	otherToMenuButton	Button
	message				string
}

var receptionStruct *ReceptionLevel

func newReceptionLevel(g *Game) {
	if len(plasmaSprites) == 0 {
		receptionStruct = &ReceptionLevel{
			protoPlasmaBg:  newStillImage("PlasmaBg.png", newRect(0, 0, 1250, 750)),
			plasmaBg:       newParallax("ParallaxPlasma.png", newRect(100, 100, 1250, 750), 4),
			plasmaMembrane: newParallax("plasmaMembrane.png", newRect(100, 300, 1250, 750), 2),

			receptorA: newReceptor("inact_receptorA.png", newRect(50, 400, 100, 100), "receptorA"),
			receptorB: newReceptor("inact_receptorB.png", newRect(350, 400, 100, 100), "receptorB"),
			receptorC: newReceptor("inact_receptorC.png", newRect(650, 400, 100, 100), "receptorC"),
			receptorD: newReceptor("inact_receptorD.png", newRect(950, 400, 100, 100), "receptorD"),

			temp_tk1A: newKinase("inact_TK1.png", newRect(50, 600, 150, 150), "temp_tk1A"),
			temp_tk1B: newKinase("inact_TK1.png", newRect(350, 600, 150, 150), "temp_tk1B"),
			temp_tk1C: newKinase("inact_TK1.png", newRect(650, 600, 150, 150), "temp_tk1C"),
			temp_tk1D: newKinase("inact_TK1.png", newRect(950, 600, 150, 150), "temp_tk1D"),
			message:	"WELCOME TO THE PLASMA MEMBRANE!\n" +
			"Drag the signal to the matching receptor\nto enter the cell!",
		}
		seedSignal = rand.Intn(4) + 1
		receptionStruct.infoButton = infoButton
		receptionStruct.otherToMenuButton = otherToMenuButton
		switch seedSignal {
		case 1:
			receptionStruct.signal = newSignal("signalA.png", newRect(500, 100, 100, 100))
			receptionStruct.signal.signalType = "signalA"
			template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
		case 2:
			receptionStruct.signal = newSignal("signalB.png", newRect(500, 100, 100, 100))
			receptionStruct.signal.signalType = "signalB"
			template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		case 3:
			receptionStruct.signal = newSignal("signalC.png", newRect(500, 100, 100, 100))
			receptionStruct.signal.signalType = "signalC"
			template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
		case 4:
			receptionStruct.signal = newSignal("signalD.png", newRect(500, 100, 100, 100))
			receptionStruct.signal.signalType = "signalD"
			template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
		}

		plasmaSprites = []GUI{
			&receptionStruct.protoPlasmaBg, &receptionStruct.plasmaBg, &receptionStruct.plasmaMembrane,
			&receptionStruct.signal, &receptionStruct.receptorA, &receptionStruct.receptorB,
			&receptionStruct.receptorC, &receptionStruct.receptorD, &receptionStruct.temp_tk1A,
			&receptionStruct.temp_tk1B, &receptionStruct.temp_tk1C, &receptionStruct.temp_tk1D,
			&receptionStruct.infoButton, &receptionStruct.otherToMenuButton,	
		}
	}
	g.stateMachine.state = receptionStruct
}

func (r *ReceptionLevel) Init(g *Game) {
	ebiten.SetWindowTitle("Cell Signaling Pathway - Signal Reception")
	state_array = plasmaSprites
}

func (r *ReceptionLevel) Update(g *Game) {
	for _, element := range plasmaSprites {
		element.update(g)
	}
	if r.receptorA.is_touching_signal {
		if matchSR(r.signal.signalType, r.receptorA.receptorType) {
			r.receptorA.animate("act_receptorA.png")
			r.signal.bind(r.receptorA)
			r.temp_tk1A.activate()
		}
	}
	if r.receptorB.is_touching_signal {
		if matchSR(r.signal.signalType, r.receptorB.receptorType) {
			r.receptorB.animate("act_receptorB.png")
			r.signal.bind(r.receptorB)
			r.temp_tk1B.activate()
		}
	}
	if r.receptorC.is_touching_signal {
		if matchSR(r.signal.signalType, r.receptorC.receptorType) {
			r.receptorC.animate("act_receptorC.png")
			r.signal.bind(r.receptorC)
			r.temp_tk1C.activate()
		}
	}
	if r.receptorD.is_touching_signal {
		if matchSR(r.signal.signalType, r.receptorD.receptorType) {
			r.receptorD.animate("act_receptorD.png")
			r.signal.bind(r.receptorD)
			r.temp_tk1D.activate()
		}
	}
	if r.temp_tk1A.rect.pos.y >= screenHeight || r.temp_tk1B.rect.pos.y >= screenHeight || r.temp_tk1C.rect.pos.y >= screenHeight || r.temp_tk1D.rect.pos.y >= screenHeight {
		ToCyto1(g)
	}
}

func (r *ReceptionLevel) Draw(g *Game, screen *ebiten.Image) {
	for _, element := range plasmaSprites {
		element.draw(screen)
	}
	defaultFont.drawFont(screen, r.message, 100, 50, color.RGBA{220, 100, 100, 50})

}
