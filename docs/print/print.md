# **Package print**

Le package print fournit une structure et des fonctions pour afficher des tableaux formatés dans la sortie standard ou dans un flux personnalisé.

---

## **Structures**

### **`Writer`**

La structure `Writer` est utilisée pour gérer l'écriture de tableaux formatés.

#### **Champs :**

- `header []string` : Contient les en-têtes des colonnes du tableau.
- `opts writerOps` : Options de configuration pour le Writer.
- `tw *tabwriter.Writer` : Instance de `tabwriter.Writer` utilisée pour formater les tableaux.

---

### **`writerOps`**

Structure interne utilisée pour configurer les options du `Writer`.

#### **Champs :**

- `output io.Writer` : Destination de la sortie (par défaut, `os.Stdout`).

---

### **`WriterOption`**

Type fonctionnel utilisé pour appliquer des options de configuration au `Writer`.

```go
type WriterOption func(w *writerOps)
```

---

## **Fonctions**

### **`New`**

Crée une nouvelle instance de `Writer`.

```go
func New(opts ...WriterOption) Writer
```

#### **Paramètres :**

- `opts ...WriterOption` : Liste d'options pour configurer le Writer.

#### **Retourne :**

- Une instance de `Writer`.

#### **Exemple :**

```go
w := print.New(print.WithOutput(os.Stdout))
```

---

### **`WithOutput`**

Définit la destination de la sortie pour le Writer.

```go
func WithOutput(output io.Writer) WriterOption
```

#### **Paramètres :**

- `output io.Writer` : Destination de la sortie (par exemple, `os.Stdout`, un fichier, etc.).

#### **Retourne :**

- Une fonction de type `WriterOption` qui configure la sortie.

#### **Exemple :**

```go
w := print.New(print.WithOutput(os.Stderr))
```

---

### **`SetHeader`**

Définit les en-têtes des colonnes pour le tableau.

```go
func (w *Writer) SetHeader(fields ...any)
```

#### **Paramètres :**

- `fields ...any` : Liste des en-têtes de colonnes (doivent être des chaînes de caractères).

#### **Comportement :**

- Convertit chaque en-tête en majuscules.
- Ajoute les en-têtes au Writer.
- Panique si un champ n'est pas une chaîne de caractères.

#### **Exemple :**

```go
w.SetHeader("Name", "Age", "Country")
```

---

### **`AddFields`**

Ajoute une ligne de données au tableau.

```go
func (w Writer) AddFields(fields ...any)
```

#### **Paramètres :**

- `fields ...any` : Liste des valeurs pour une ligne (doivent correspondre au nombre d'en-têtes).

#### **Comportement :**

- Formate chaque champ avant de l'ajouter.
- Panique si le nombre de champs ne correspond pas au nombre d'en-têtes.

#### **Exemple :**

```go
w.AddFields("Alice", 30, "France")
w.AddFields("Bob", 25, "USA")
```

---

### **`PrintTable`**

Affiche le tableau formaté dans la sortie configurée.

```go
func (w Writer) PrintTable()
```

#### **Comportement :**

- Vide le tampon interne et affiche le tableau.

#### **Exemple :**

```go
w.PrintTable()
```

---

## **Exemple d'utilisation**

Voici un exemple complet d'utilisation du package print :

```go
package main

import (
	"os"

	"github.com/orange-cloudavenue/common-go/print"
)

func main() {
	// Créer un Writer avec une sortie personnalisée
	w := print.New(print.WithOutput(os.Stdout))

	// Définir les en-têtes
	w.SetHeader("Name", "Age", "Country")

	// Ajouter des lignes
	w.AddFields("Alice", 30, "France")
	w.AddFields("Bob", 25, "USA")
	w.AddFields("Charlie", 35, "UK")

	// Afficher le tableau
	w.PrintTable()
}
```

---

## **Détails supplémentaires**

### **Formatage des champs**

Les champs ajoutés au tableau sont formatés avant d'être écrits. Le formatage est géré par une fonction interne `fieldFormat`.

### **Formatage des lignes**

Les lignes sont formatées à l'aide de la fonction interne `format`, qui applique un espacement uniforme entre les colonnes.

---

## **Licence**

Ce fichier est distribué sous la licence **Mozilla Public License 2.0**.  
Pour plus d'informations, consultez le fichier [LICENSE](https://www.mozilla.org/en-US/MPL/2.0/).
