Voici la documentation générée au format Markdown pour le fichier urn.go :

---

# **Package validators**

Le package validators fournit des validateurs personnalisés pour vérifier la conformité des champs dans les structures Go. Ce fichier contient le validateur `URN`, qui vérifie si une chaîne de caractères est une URN (Uniform Resource Name) valide.

---

## **Validateur `URN`**

Le validateur `URN` est utilisé pour vérifier si une chaîne de caractères est une URN valide en fonction d'un type URN spécifique.

### **Définition**

```go
var URN = &CustomValidator{
	Key: "urn",
	Func: func(fl validator.FieldLevel) bool {
		u, err := urn.FindURNTypeFromString(fl.Param())
		if err != nil {
			return false
		}

		return strings.Contains(fl.Field().String(), u.String())
	},
}
```

---

### **Fonctionnalité**

Le validateur `URN` :
1. Utilise la fonction `urn.FindURNTypeFromString` pour déterminer le type de la URN à partir du paramètre fourni dans le tag `urn`.
2. Vérifie si la valeur du champ contient le type URN attendu.
3. Retourne `true` si la validation réussit, sinon `false`.

---

### **Utilisation**

#### **Exemple de structure avec le tag `urn` :**

```go
type ExampleStruct struct {
	Resource string `validate:"urn=example-type"` // Vérifie si la chaîne est une URN valide du type "example-type"
}
```

#### **Exemple d'application du validateur :**

```go
package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"your_project_path/validators"
)

func main() {
	validate := validator.New()

	// Enregistrer le validateur personnalisé "urn"
	validate.RegisterValidation(validators.URN.Key, validators.URN.Func)

	// Exemple de structure
	example := ExampleStruct{
		Resource: "urn:example-type:12345",
	}

	// Valider la structure
	err := validate.Struct(&example)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return
	}

	fmt.Println("Validation passed!")
}
```

---

### **Comportement**

1. **Cas valide :**
   - Si le champ contient une URN valide correspondant au type spécifié, la validation réussit.
   - Exemple : Pour `urn=example-type`, la valeur `urn:example-type:12345` est valide.

2. **Cas invalide :**
   - Si le champ ne contient pas une URN valide ou si le type URN ne correspond pas, la validation échoue.
   - Exemple : Pour `urn=example-type`, la valeur `urn:other-type:12345` est invalide.

---

### **Dépendances**

- **`urn.FindURNTypeFromString`** : Fonction utilisée pour déterminer le type URN à partir d'une chaîne.
- **`strings.Contains`** : Vérifie si la valeur du champ contient le type URN attendu.

---

### **Exemple de comportement**

#### **Cas 1 : Validation réussie**
```go
example := ExampleStruct{
	Resource: "urn:example-type:12345",
}
```
Sortie :
```plaintext
Validation passed!
```

#### **Cas 2 : Validation échouée**
```go
example := ExampleStruct{
	Resource: "urn:other-type:12345",
}
```
Sortie :
```plaintext
Validation error: Key: 'ExampleStruct.Resource' Error:Field validation for 'Resource' failed on the 'urn' tag
```

---

## **Licence**

Ce fichier est distribué sous la licence **Mozilla Public License 2.0**.  
Pour plus d'informations, consultez le fichier [LICENSE](https://www.mozilla.org/en-US/MPL/2.0/).

---

## **Résumé**

Le validateur `URN` est un outil pratique pour valider les URN dans les structures Go. Il garantit que les champs respectent un format URN spécifique, ce qui est utile pour les applications nécessitant une gestion stricte des identifiants de ressources.