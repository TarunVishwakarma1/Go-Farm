package farm

type Animal interface {
	Speak() string
	NumberOfLegs() int
	GetType() string
	GetName() string
	SetID(id int)
}

type Cow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Chicken struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type AnimalReponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (c *Cow) Speak() string     { return "Moo!" }
func (c *Cow) NumberOfLegs() int { return 4 }
func (c *Cow) GetType() string   { return "Cow" }
func (c *Cow) GetName() string   { return c.Name }
func (c *Cow) SetID(id int)      { c.ID = id }

func (c *Chicken) Speak() string     { return "Cluck!" }
func (c *Chicken) NumberOfLegs() int { return 2 }
func (c *Chicken) GetType() string   { return "Chicken" }
func (c *Chicken) GetName() string   { return c.Name }
func (c *Chicken) SetID(id int)      { c.ID = id }

func (c *AnimalReponse) Speak() string     { return "Cluck!" }
func (c *AnimalReponse) NumberOfLegs() int { return 2 }
func (c *AnimalReponse) GetType() string   { return "Chicken" }
func (c *AnimalReponse) GetName() string   { return c.Name }
func (c *AnimalReponse) SetID(id int)      { c.ID = id }
