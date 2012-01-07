package main

import (
   "fmt"
   //"strconv"
)

type Space struct {
   color int
   freedom int
}

type Game struct {
   board [][]Space
   w, h int
   s1, s2 int
}

func (p *Game) Init(w, h int) *Game {
   p.w, p.h = w, h
   p.board = make([][]Space, h)
   for row := 0; row < h; row++ {
      p.board[row] = make([]Space, w)
      for col := 0; col < w; col++ {
         freedoms := 4
         if 0 == col || (w - 1) == col {
            freedoms--
         }
         if 0 == row || (h - 1) == row {
            freedoms--
         }
         p.board[row][col].color = 0
         p.board[row][col].freedom = freedoms
      }
   }
   return p
}

func (p *Game) Print() {
   f := func(s Space) (o string) {
      switch s.color {
      case -1:
         o = "#"
      case 0:
         o = "." //strconv.Itoa(s.freedom)
      case 1:
         o = "X"
      case 2:
         o = "O"
      }
      return o
   }
   for r := 0; r < p.h; r++ {
      for c := 0; c < p.w; c++ {
         print(f(p.board[r][c]))
      }
      println()
   }
}

func (p *Game) ValidMove(r, c int) bool {
   return p.InBoundary(r, c) && p.board[r][c].color == 0
}

func (p *Game) InBoundary(r, c int) bool {
   return !((r < 0 || r >= p.h) || (c < 0 || c >= p.w))
}

func (p *Game) GetMove(msg string) (r, c int) {
   r, c = -1, -1
   for !p.ValidMove(r, c) {
      fmt.Println(msg, "Please enter a move: (int, int)")
      fmt.Scanf("%d,%d", &r, &c)
   }
   return r, c
}

// Assumes we are given a valid move.
func (p *Game) PlaceColor(r, c, color int) {
   p.board[r][c].color = color
   p.AdjustFreedoms(r, c, -1)
}

func (p *Game) RemoveColor(r, c int) {
   if p.board[r][c].color == 2 {
      p.s1++
   }
   if p.board[r][c].color == 1 {
      p.s2++
   }
   p.board[r][c].color = 0
   p.AdjustFreedoms(r, c, 1)
}

func (p *Game) AdjustFreedoms(r, c, v int) {
   p.SetFreedom(r - 1, c, v)
   p.SetFreedom(r + 1, c, v)
   p.SetFreedom(r, c - 1, v)
   p.SetFreedom(r, c + 1, v)
}

func (p *Game) SetFreedom(r, c, freedom int) {
   if p.InBoundary(r, c) {
      p.board[r][c].freedom += freedom
   }
}


func (p *Game) Capture() {
   for r := 0; r < p.h; r++ {
      for c := 0; c < p.w; c++ {
         if p.board[r][c].freedom < 1 {
            defer p.RemoveColor(r, c)
         }
      }
   }
}

func (p *Game) ResolveTurn(r1, c1, r2, c2 int) {
   if r1 == r2 && c1 == c2 {
      p.PlaceColor(r1, c1, -1)
   } else {
      p.PlaceColor(r1, c1, 1)
      p.PlaceColor(r2, c2, 2)
   }
   p.Capture()
}

func (p *Game) DoTurn() {
   p.Print()
   r1, c1 := p.GetMove("X")
   r2, c2 := p.GetMove("O")
   p.ResolveTurn(r1, c1, r2, c2)
}

func main() {
   game := new(Game).Init(5, 5)
   for {
      game.DoTurn()
   }
}
