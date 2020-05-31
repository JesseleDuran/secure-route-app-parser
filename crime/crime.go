package crime

import (
  "log"
  "os"
  "proyecto/text"
  "time"

  "github.com/gocarina/gocsv"
)

type Crime struct {
  Day             string
  Date            string
  Hour            string
  Department      string
  Municipality    string
  Neighborhood    string
  Type            string
  Value           int
  VictimGender    string
  VictimTransport string
  VictimAge       int
  AggressorWeapon string
}

type Crimes []Crime

// FromCSV given a path of a csv retrieve a collection of crimes.
func FromCSV(path string) Crimes {
  type csv struct {
    Day             string `csv:"Día"`
    Date            string `csv:"Fecha"`
    Hour            string `csv:"Hora"`
    Department      string `csv:"Departamento"`
    Municipality    string `csv:"Municipio"`
    Neighborhood    string `csv:"Barrio"`
    VictimGender    string `csv:"Sexo"`
    VictimTransport string `csv:"Móvil Victima"`
    VictimAge       int    `csv:"Edad"`
    AggressorWeapon string `csv:"Arma empleada"`
  }

  //get csv info.
  var cc []*csv
  f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  if err := gocsv.UnmarshalFile(f, &cc); err != nil {
    panic(err)
  }

  //parse result.
  result := make(Crimes, 0)
  for i := range cc {
    date, err := time.Parse("1/02/2006 3:04:05 PM", cc[i].Date)
    if err != nil {
      log.Println("date parse", err)
    }
    hour, err := time.Parse("1/02/2006 03:04:05 PM", cc[i].Hour)
    if err != nil {
      log.Println("hour parse:", err)
    }
    result = append(result, Crime{
      Day:             cc[i].Day,
      Date:            date.Format("02/01/2006"),
      Hour:            hour.Format("15:04:05"),
      Department:      text.Normalize(cc[i].Department),
      Municipality:    text.Normalize(cc[i].Municipality),
      Neighborhood:    text.Normalize(cc[i].Neighborhood),
      Type:            "hurto", //TODO: this should be passed to the function.
      Value:           1,       //TODO: this should depend of the judgment.
      VictimGender:    text.Normalize(cc[i].VictimGender),
      VictimTransport: text.Normalize(cc[i].VictimTransport),
      VictimAge:       cc[i].VictimAge,
      AggressorWeapon: text.Normalize(cc[i].AggressorWeapon),
    })
  }
  return result
}
