package main

import (
   "fmt"
   "strconv"
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
         o = strconv.Itoa(s.freedom)
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
   return p.InBoundary(r, c) && p.spaces[r][c].color == 0 && p.spaces[r][c].freedom != 0
}

func (p *Board) InBoundary(r, c int) bool {
   return !((r < 0 || r >= p.h) || (c < 0 || c >= p.w))
}

func (p *Board) GetMove() (r, c int) {
   r, c = -1, -1
   for !p.ValidMove(r, c) {
      fmt.Println("Please enter a move: (int, int)")
      fmt.Scanf("%d,%d", &r, &c)
   }
   return r, c
}

// Assumes we are given a valid move.
func (p *Board) PlaceColor(r, c, color int) {
   p.spaces[r][c].color = color
   p.AdjustFreedoms(r, c, color)
}

func (p *Board) AdjustFreedoms(r, c, color int) {
   p.SetFreedom(r - 1, c, -1)
   p.SetFreedom(r + 1, c, -1)
   p.SetFreedom(r, c - 1, -1)
   p.SetFreedom(r, c + 1, -1)
}

func (p *Board) SetFreedom(r, c, freedom int) {
   if p.InBoundary(r, c) {
      p.spaces[r][c].freedom += freedom
   }
}

func main() {
   board := new(Board).Init(10, 10)
   board.PlaceColor(5, 5, 1)
   board.Print()
   r, c := board.GetMove()
   board.PlaceColor(r, c, 1)
   fmt.Printf("Got two numbers %d, %d\n", r, c)
   board.Print()
}
