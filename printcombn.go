package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i <= 9; i++ {
			z01.PrintRune(rune(48 + i))
			if i == 9 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}

		}
	}
	if n == 2 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				if i < j {
					z01.PrintRune(rune(48 + i))
					z01.PrintRune(rune(48 + j))
					if i == 8 {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	if n == 3 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					if i < j && j < k {
						z01.PrintRune(rune(48 + i))
						z01.PrintRune(rune(48 + j))
						z01.PrintRune(rune(48 + k))
						if i == 7 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	if n == 4 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						if i < j && j < k && k < l {
							z01.PrintRune(rune(48 + i))
							z01.PrintRune(rune(48 + j))
							z01.PrintRune(rune(48 + k))
							z01.PrintRune(rune(48 + l))
							if i == 6 {
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
	if n == 5 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for m := 0; m <= 9; m++ {
							if i < j && j < k && k < l && l < m {
								z01.PrintRune(rune(48 + i))
								z01.PrintRune(rune(48 + j))
								z01.PrintRune(rune(48 + k))
								z01.PrintRune(rune(48 + l))
								z01.PrintRune(rune(48 + m))
								if i == 5 {
									z01.PrintRune('\n')
								} else {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 6 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for m := 0; m <= 9; m++ {
							for n := 0; n <= 9; n++ {
								if i < j && j < k && k < l && l < m && m < n {
									z01.PrintRune(rune(48 + i))
									z01.PrintRune(rune(48 + j))
									z01.PrintRune(rune(48 + k))
									z01.PrintRune(rune(48 + l))
									z01.PrintRune(rune(48 + m))
									z01.PrintRune(rune(48 + n))
									if i == 4 {
										z01.PrintRune('\n')
									} else {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 7 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for m := 0; m <= 9; m++ {
							for n := 0; n <= 9; n++ {
								for o := 0; o <= 9; o++ {
									if i < j && j < k && k < l && l < m && m < n && n < o {
										z01.PrintRune(rune(48 + i))
										z01.PrintRune(rune(48 + j))
										z01.PrintRune(rune(48 + k))
										z01.PrintRune(rune(48 + l))
										z01.PrintRune(rune(48 + m))
										z01.PrintRune(rune(48 + n))
										z01.PrintRune(rune(48 + o))
										if i == 3 {
											z01.PrintRune('\n')
										} else {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 8 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for m := 0; m <= 9; m++ {
							for n := 0; n <= 9; n++ {
								for o := 0; o <= 9; o++ {
									for p := 0; p <= 9; p++ {
										if i < j && j < k && k < l && l < m && m < n && n < o && o < p {
											z01.PrintRune(rune(48 + i))
											z01.PrintRune(rune(48 + j))
											z01.PrintRune(rune(48 + k))
											z01.PrintRune(rune(48 + l))
											z01.PrintRune(rune(48 + m))
											z01.PrintRune(rune(48 + n))
											z01.PrintRune(rune(48 + o))
											z01.PrintRune(rune(48 + p))
											if i == 2 {
												z01.PrintRune('\n')
											} else {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	if n == 9 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for m := 0; m <= 9; m++ {
							for n := 0; n <= 9; n++ {
								for o := 0; o <= 9; o++ {
									for p := 0; p <= 9; p++ {
										for q := 0; q <= 9; q++ {
											if i < j && j < k && k < l && l < m && m < n && n < o && o < p && p < q {
												z01.PrintRune(rune(48 + i))
												z01.PrintRune(rune(48 + j))
												z01.PrintRune(rune(48 + k))
												z01.PrintRune(rune(48 + l))
												z01.PrintRune(rune(48 + m))
												z01.PrintRune(rune(48 + n))
												z01.PrintRune(rune(48 + o))
												z01.PrintRune(rune(48 + p))
												z01.PrintRune(rune(48 + q))
												if i == 1 {
													z01.PrintRune('\n')
												} else {
													z01.PrintRune(',')
													z01.PrintRune(' ')
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
