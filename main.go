package main

import (
	"fmt"
	"math/rand"

	task1 "github.com/artsadert/golang_exam/task_1"
	task2 "github.com/artsadert/golang_exam/task_2"
	task3 "github.com/artsadert/golang_exam/task_3"
)

func main() {
	// task_1
	fmt.Println()
	fmt.Println("             task1")
	fmt.Println()

	PepeSchenele1 := task1.NewPepeSchenele(1, 2, 3)
	PepeSchenele2 := task1.NewPepeSchenele(2, 2, 3)

	if PepeSchenele1.GetRaiting() == PepeSchenele2.GetRaiting() {
		fmt.Println("PepeSchenele1 is equal PepeSchenele2")
	} else if PepeSchenele1.GetRaiting() < PepeSchenele2.GetRaiting() {
		fmt.Println("PepeSchenele2 is better than PepeSchenele1")
	} else {
		fmt.Println("PepeSchenele1 is better than PepeSchenele2")
	}

	// task_2
	fmt.Println()
	fmt.Println("             task2")
	fmt.Println()

	memes := make([]task2.BrainrotMeme, 7)
	for i := 0; i < 7; i++ {
		meme, err := task2.NewBrainrotMeme(fmt.Sprintf("meme_%d", i), i+2, "Sigma", rand.Float64()*100)
		if err != nil {
			fmt.Println("cannot create meme")
		}

		memes[i] = *meme
	}

	fmt.Println(task2.FindTopTrending(memes, 50))
	fmt.Println(task2.FilterByComplexCondition(memes))
	fmt.Println(task2.CalculateCategoryImpact(memes))

	// task_3
	fmt.Println()
	fmt.Println("             task3")
	fmt.Println()

	phrase := "The quick brown fox jumps over the lazy dog"
	fmt.Println(task3.EncruptPhrase(phrase))
}
