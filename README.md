### PascalCase und camelCase
- Pascalcase (public), camelCase (private)

### func Definition 
- fun (p *WelcomeRepo) GetWelcomeMessage() string {..}
- (p *WelcomeRepo) = receiver, GetWelcomeMessage() = name, string Rückgabewert
- Receiver brauchen wir, da man ein Interface mit mehreren gleichen Methoden haben kann (java this. in GO explizit davor schreiben)
- Receiver Pointer brauchen wir auch, sonst würden wir mitt jedem Aufruf eine komplette Kopie des Structs erstellen