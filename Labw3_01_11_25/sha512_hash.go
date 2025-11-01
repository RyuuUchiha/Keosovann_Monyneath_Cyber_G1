package main

import "Labw3_01_11_25/utils/crack"

func lab3() {

	defaultTarget := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"
	defaultVerbose := "results/verbose3.txt"

	cracker("SHA512", crack.SHA512Hash, defaultTarget, defaultVerbose)
}
