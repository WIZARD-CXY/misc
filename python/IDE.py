import wx,os

def load(event):
	file = open(filename.GetValue())
	contents.SetValue(file.read())
	file.close()

def save_run(event):
	file=open(filename.GetValue(),'w')
	file.write(contents.GetValue())
	file.close()
        os.system("gcc "+filename.GetValue())
	os.system("./a.out")

app = wx.App()
win = wx.Frame(None,title="My small C IDE ",size=(410,335))

bkg = wx.Panel(win)

loadButton = wx.Button(bkg, label='Open')
loadButton.Bind(wx.EVT_BUTTON,load)

saveButton = wx.Button(bkg, label='Save_Run')
saveButton.Bind(wx.EVT_BUTTON,save_run)

transferButton = wx.Button(bkg, label='Transfer')

filename=wx.TextCtrl(bkg)
contents=wx.TextCtrl(bkg,style=wx.TE_MULTILINE | wx.HSCROLL)

hbox=wx.BoxSizer()
hbox.Add(filename,proportion=1,flag=wx.EXPAND)
hbox.Add(loadButton,proportion=0,flag=wx.LEFT,border=5)
hbox.Add(saveButton,proportion=0,flag=wx.LEFT,border=5)
hbox.Add(transferButton,proportion=0,flag=wx.LEFT,border=5)

vbox=wx.BoxSizer(wx.VERTICAL)
vbox.Add(hbox,proportion=0,flag=wx.EXPAND | wx.ALL,border=5)
vbox.Add(contents,proportion=1,
		flag=wx.EXPAND | wx.LEFT | wx.BOTTOM | wx.RIGHT  ,border=5)

bkg.SetSizer(vbox)

win.Show()
app.MainLoop()
