package main

import (
	"fmt"
	"math/rand"
)

func ToPlasma(g *Game) {
	g.receptionSprites = nil

	scene = "Signal Reception"
	g.stateMachine.changeState(g, scene)
}

func ToMenu(g *Game) {
	// Set length of sprite array to 0
	g.menuSprites = nil

	g.reset()
	scene = "Main Menu"
	g.stateMachine.changeState(g, scene)
}

func ToCyto1(g *Game) {
	g.transductionSprites = nil

	scene = "Signal Transduction"
	g.stateMachine.changeState(g, scene)
}

func ToNucleus(g *Game) {
	g.transcriptionSprites = nil

	scene = "Transcription"
	g.stateMachine.changeState(g, scene)
}

func ToCyto2(g *Game) {
	g.translationSprites = nil

	scene = "Translation"
	g.stateMachine.changeState(g, scene)
}

func ToLevelSelect(g *Game) {
	g.levSelSprites = nil

	scene = "Level Selection"
	g.stateMachine.changeState(g, scene)
}

func ToAbout(g *Game) {
	g.aboutSprites = nil

	scene = "About"
	g.stateMachine.changeState(g, scene)
}

func (g *Game) reset() {
	// Set seed signal to random integer
	seedSignal = rand.Intn(4) + 1

	// Set template to random codons
	switch seedSignal {
	case 1:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ACT"}
	case 2:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	case 3:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATC"}
	case 4:
		template = [5]string{"TAC", randomDNACodon(), randomDNACodon(), randomDNACodon(), "ATT"}
	}

	// Set dna, rna, and proteins to random codons
	for x := 0; x < 5; x++ {
		dna[x] = newTemplate("DNA.png", newRect(200*x, 400, 150, 150), template[x], x)
	}
	for x := 0; x < 5; x++ {
		rna[x] = newTranscript("RNA"+fmt.Sprint(x)+".png", newRect((100*x)-0, 0, 150, 150), transcribe(template[x]), true)
	}

	for x := 0; x < 5; x++ {
		mrna[x] = newTemplate("DNA.png", newRect(100*x, 250, 150, 150), transcribe(dna[x].codon), x)
	}
	for x := 0; x < 5; x++ {
		protein[x] = newTranscript("aminoAcid.png", newRect(125+(150*x), 225, 150, 150), translate(mrna[x].codon), false)
	}
}
