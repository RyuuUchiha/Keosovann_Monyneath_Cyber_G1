package main

import "Labw3_01_11_25/utils/crack"

func lab2() {

	defaultTarget := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"
	defaultVerbose := "results/verbose2.txt"

	cracker("SHA1", crack.SHA1Hash, defaultTarget, defaultVerbose)
}
