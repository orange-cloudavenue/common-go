# **Package validators**

Le package validators fournit des validateurs personnalisés pour vérifier la conformité des champs dans les structures Go. Ce fichier contient des validateurs pour interdire les majuscules et les espaces dans les chaînes de caractères.

---

## **Validateurs**

### **`DisallowUpper`**

Le validateur `DisallowUpper` vérifie qu'une chaîne de caractères ne contient pas de lettres majuscules.

#### **Définition**

```go
var DisallowUpper = &CustomValidator{
	Key: "disallow_upper",
	Func: func(fl validator.FieldLevel) bool {
		for _, r := range fl.Field().String() {
			if unicode.IsUpper(r) {
				return false
			}
		}
		return true
	},
}
```

#### **Fonctionnalité**
- Parcourt chaque caractère de la chaîne.
- Retourne `false` si un caractère majuscule est trouvé.
- Retourne `true` si aucun caractère majuscule n'est présent.

#### **Exemple d'utilisation**

```go
type ExampleStruct struct {
	Username string `validate:"disallow_upper"` // Vérifie que la chaîne ne contient pas de majuscules
}
```

---

### **`DisallowSpace`**

Le validateur `DisallowSpace` vérifie qu'une chaîne de caractères ne contient pas d'espaces.

#### **Définition**

```go
var DisallowSpace = &CustomValidator{
	Key: "disallow_space",
	Func: func(fl validator.FieldLevel) bool {
		for _, r := range fl.Field().String() {
			if unicode.IsSpace(r) {
				return false
			}
		}
		return true
	},
}
```

#### **Fonctionnalité**
- Parcourt chaque caractère de la chaîne.
- Retourne `false` si un espace est trouvé.
- Retourne `true` si aucun espace n'est présent.

#### **Exemple d'utilisation**

```go
type ExampleStruct struct {
	Password string `validate:"disallow_space"` // Vérifie que la chaîne ne contient pas d'espaces
}
```

---

## **Exemple d'utilisation globale**

Voici un exemple complet d'utilisation des validateurs dans une structure Go :

```go
package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"your_project_path/validators"
)

type ExampleStruct struct {
	Username string `validate:"disallow_upper"`
	Password string `validate:"disallow_space"`
}

func main() {
	validate := validator.New()

	// Enregistrer les validateurs personnalisés
	validate.RegisterValidation(validators.DisallowUpper.Key, validators.DisallowUpper.Func)
	validate.RegisterValidation(validators.DisallowSpace.Key, validators.DisallowSpace.Func)

	// Exemple de structure
	example := ExampleStruct{
		Username: "john_doe",
		Password: "mypassword",
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

## **Comportement**

### **`DisallowUpper`**

1. **Cas valide :**
   - La chaîne ne contient pas de majuscules.
   - Exemple : `"john_doe"`.

2. **Cas invalide :**
   - La chaîne contient une ou plusieurs majuscules.
   - Exemple : `"John_Doe"`.

### **`DisallowSpace`**

1. **Cas valide :**
   - La chaîne ne contient pas d'espaces.
   - Exemple : `"mypassword"`.

2. **Cas invalide :**
   - La chaîne contient un ou plusieurs espaces.
   - Exemple : `"my password"`.

---

## **Licence**

Ce fichier est distribué sous la licence **Mozilla Public License 2.0**.  
Pour plus d'informations, consultez le fichier [LICENSE](https://www.mozilla.org/en-US/MPL/2.0/).

---

## **Résumé**

Le fichier cases.go fournit deux validateurs personnalisés :
- **`DisallowUpper`** : Vérifie qu'une chaîne ne contient pas de majuscules.
- **`DisallowSpace`** : Vérifie qu'une chaîne ne contient pas d'espaces.

Ces validateurs sont utiles pour appliquer des règles strictes sur les chaînes de caractères dans les structures Go.