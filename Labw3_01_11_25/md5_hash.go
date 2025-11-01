package main

import "Labw3_01_11_25/utils/crack"

func lab1() {

	defaultTarget := "6a85dfd77d9cb35770c9dc6728d73d3f"
	defaultVerbose := "results/verbose1.txt"

	cracker("MD5", crack.MD5Hash, defaultTarget, defaultVerbose)
}
