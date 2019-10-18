package module

type Module interface {
    SetUp()
}

type module struct {
    mi  Module
}

var mods []*module

func Register(mi ...Module) {
   for _,m := range mi{
       mod := new(module)
       mod.mi = m
       mods = append(mods, mod)
   }
}

func Init(){
    for _,m := range mods{
        m.mi.SetUp()
    }
}
