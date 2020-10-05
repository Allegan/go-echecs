module github.com/Allegan/go-echecs

go 1.15

replace github.com/Allegan/go-echecs/pkg/board => ./pkg/board
replace github.com/Allegan/go-echecs/pkg/piece => ./pkg/piece
replace github.com/Allegan/go-echecs/pkg/color => ./pkg/color

require github.com/sirupsen/logrus v1.7.0
