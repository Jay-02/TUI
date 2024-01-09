package main
import(
"fmt"
"os"
tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
)
type status int
const (
todo status = iota
inProgress
done
)
type Task struct{
status status
title string
description string
}
func (t Task) FilterValue() string{
    return t.title
    }
func (t Task) Title() string {
    return t.title
    }

func (t Task) Description() string {
return t.description
}
type Model struct{
list list.Model
err error
}
func New() *Model{
    return &Model{}
}
func (m *Model) initList(width, height int){
    m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
    m.list.Title = "To Do"
    m.list.SetItems([]list.Item{
    Task{ status: todo, title: "meow", description: "cat"},
    Task{ status: inProgress, title: "woof", description: "dog"},

        })
}
func (m Model) Init() tea.Cmd{
    return nil
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
    switch msg:=msg.(type){
    case tea.WindowSizeMsg:
        m.initList(msg.Width, msg.Height)
    }
    var cmd tea.Cmd
    m.list, cmd = m.list.Update(msg)
     return m, cmd
}
func (m Model) View() string{
    return m.list.View()
}
func main(){
    m:=New()
    p:=tea.NewProgram(m)
    if err:=p.Start();err !=nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
