# **Validateur `default`**

## **Utilisation**

### **Exemple de structure avec des tags `default` :**

```go
type ExampleStruct struct {
	Name  string  `validate:"default=John Doe"` // Définit "John Doe" comme valeur par défaut
	Age   int     `validate:"default=30"`      // Définit 30 comme valeur par défaut
	Admin bool    `validate:"default=true"`    // Définit true comme valeur par défaut
	Score float64 `validate:"default=99.5"`    // Définit 99.5 comme valeur par défaut
}
```

### **Exemple d'application du validateur :**

```go
package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"your_project_path/validators"
)

func main() {
	validate := validator.New()

	// Enregistrer le validateur personnalisé "default"
	validate.RegisterValidation(validators.Default.Key, validators.Default.Func)

	// Exemple de structure
	example := ExampleStruct{}

	// Appliquer les valeurs par défaut
	err := validate.Struct(&example)
	if err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return
	}

	// Afficher la structure avec les valeurs par défaut appliquées
	fmt.Printf("Struct with defaults applied: %+v\n", example)
}
```

---

## **Comportement par type**

| **Type**       | **Exemple de tag**       | **Valeur appliquée**          |
|-----------------|--------------------------|--------------------------------|
| `string`        | `default=John Doe`       | Définit `"John Doe"`           |
| `int`           | `default=42`            | Définit `42`                   |
| `uint`          | `default=100`           | Définit `100`                  |
| `float64`       | `default=3.14`          | Définit `3.14`                 |
| `bool`          | `default=true`          | Définit `true`                 |

---

## **Gestion des erreurs**

1. Si le champ n'est pas vide, le validateur retourne `true` sans modifier la valeur.
2. Si le champ est vide mais que la valeur par défaut est invalide (par exemple, `default=invalid` pour un champ `int`), le validateur retourne `false`.
3. Si le type du champ n'est pas pris en charge, le validateur retourne également `false`.

---

## **Limites**

- Le validateur ne prend pas en charge les types complexes comme les tableaux, les slices, les maps ou les structures imbriquées.
- Si un type non pris en charge est utilisé avec le tag `default`, le validateur échouera.

---

## **Résumé**

Le validateur `Default` est un outil pratique pour définir des valeurs par défaut dans vos structures Go. Il est particulièrement utile pour garantir que les champs critiques ont une valeur par défaut, même si l'utilisateur ne les initialise pas explicitement.