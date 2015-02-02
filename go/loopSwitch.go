loop:
for {
        switch expr {
        case foo:
                if condA {
                        doA()
                        break // like 'goto A'
                }

                if condB {
                        doB()
                        break loop // like 'goto B'                        
                }

                doC()
        case bar:
                // ...
        }
A:
        doX()
        // ...
}

B:
doY()