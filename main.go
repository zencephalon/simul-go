package main

import (
   "fmt"
)

type Space struct {
   color int
   freedom int
}

type Board struct {
   spaces [][]Space
   w, h int
}

func (p *Board) Init(w, h int) *Board {
   p.w, p.h = w, h
   p.spaces = make([][]Space, h)
   for row := 0; row < h; row++ {
      p.spaces[row] = make([]Space, w)
      for col := 0; col < w; col++ {
         freedoms := 4
         if 0 == col || (w - 1) == col {
            freedoms--
         }
         if 0 == row || (h - 1) == row {
            freedoms--
         }
         p.spaces[row][col].color = 0
         p.spaces[row][col].freedom = freedoms
      }
   }
   return p
}

func (p *Board) Print() {
   f := func(s Space) (o string) {
      switch s.color {
      case 0:
         o = "."
      case 1:
         o = "x"
      case 2:
         o = "o"
      }
      return o
   }
   for r := 0; r < p.h; r++ {
      for c := 0; c < p.w; c++ {
         print(f(p.spaces[r][c]))
      }
      println()
   }
}

func (p *Board) ValidMove(r, c int) bool {
   return !((r < 0 || r >= p.h) || (c < 0 || c >= p.w)) && p.spaces[r][c].color == 0 && p.spaces[r][c].freedom != 0
}

func GetMove() (r, c int) {
   fmt.Println("Please enter a move: ")
   fmt.Scanf("%d,%d", &r, &c)
   return r, c
}

func main() {
   board := new(Board).Init(10, 10)
   board.Print()
   r, c := GetMove()
   fmt.Printf("Got two numbers %d, %d", r, c)
}
