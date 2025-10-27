package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// SimularProofOfWork simula la búsqueda de una prueba de trabajo de blockchain.
// La dificultad determina el número de ceros iniciales que debe tener el hash.
// La complejidad crece exponencialmente con la dificultad.
// Para un computador personal, dificultad 5-6 suele tardar unos segundos.
func SimularProofOfWork(blockData string, dificultad int) (string, int) {
	targetPrefix := strings.Repeat("0", dificultad)
	nonce := 0
	for {
		data := fmt.Sprintf("%s%d", blockData, nonce)
		hashBytes := sha256.Sum256([]byte(data))
		hashString := fmt.Sprintf("%x", hashBytes)

		if strings.HasPrefix(hashString, targetPrefix) {
			return hashString, nonce
		}
		nonce++
	}
}

// EncontrarPrimos busca todos los números primos hasta un entero max.
// Utiliza un enfoque de prueba por división, cuya complejidad es alta (aprox.O(n^1.5)).
func EncontrarPrimos(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// CalcularTrazaDeProductoDeMatrices multiplica dos matrices NxN y devuelve la traza
// de la matriz resultante. La complejidad del cómputo es O(n^3).
func CalcularTrazaDeProductoDeMatrices(n int) int {
	// Se crean dos matrices con valores aleatorios para la multiplicación.
	m1 := make([][]int, n)
	m2 := make([][]int, n)
	for i := 0; i < n; i++ {
		m1[i] = make([]int, n)
		m2[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m1[i][j] = rand.Intn(10)
			m2[i][j] = rand.Intn(10)
		}
	}
	// Se realiza la multiplicación y se calcula la traza en el proceso.
	trace := 0
	for i := 0; i < n; i++ {
		sum := 0
		for k := 0; k < n; k++ {
			sum += m1[i][k] * m2[k][i]
		}
		trace += sum
	}
	return trace
}

func EjecucionSecuencial(n int, umbral int, blockData string, dificultad int, maxPrimos int) (string, time.Duration) {
	var ramaGanadora string
	inicio := time.Now()
	resultadoTraza := CalcularTrazaDeProductoDeMatrices(n)
	if resultadoTraza > umbral {
		fmt.Println("Se ejecutará la rama A")
		ramaGanadora = "A"
		SimularProofOfWork(blockData, dificultad)
	} else {
		fmt.Println("Se ejecutará la rama B")
		ramaGanadora = "B"
		EncontrarPrimos(maxPrimos)
	}
	fin := time.Now()
	tiempoFinal := fin.Sub(inicio)
	return ramaGanadora, tiempoFinal
}

func main() {
	n := flag.Int("n", 100, "Dimensión de las matrices")
	umbral := flag.Int("umbral", 10000, "Valor umbral para decidir la rama")
	archivo := flag.String("archivo", "salida.txt", "Archivo de salida")

	flag.Parse()

	npunt := *n
	umbralpunt := *umbral
	archivopunt := *archivo

	fmt.Printf("Configuración de la simulación:\n")
	fmt.Printf("Dimensión de las matrices: %d\n", npunt)
	fmt.Printf("Valor umbral: %d\n", umbralpunt)
	fmt.Printf("Archivo de salida: %s\n", archivopunt)
}
