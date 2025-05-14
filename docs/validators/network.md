# **Package validators**

Le package validators fournit des validateurs personnalisés pour vérifier la conformité des champs dans les structures Go. Ce fichier contient des validateurs pour les adresses IPv4, les ports TCP/UDP et les plages de ports.

---

## **Validateurs**

### **`IPV4Range`**

Le validateur `IPV4Range` vérifie si une chaîne de caractères représente une plage valide d'adresses IPv4.

#### **Définition**

```go
var IPV4Range = &CustomValidator{
	Key: "ipv4_range",
	Func: func(fl validator.FieldLevel) bool {
		parts := strings.Split(fl.Field().String(), "-")
		if len(parts) != 2 {
			return false
		}

		firstIP := net.ParseIP(parts[0])
		secondIP := net.ParseIP(parts[1])
		if firstIP.To4() == nil || secondIP.To4() == nil {
			return false
		}

		if bytes.Compare(firstIP.To4(), secondIP.To4()) >= 0 {
			return false
		}

		return true
	},
}
```

#### **Fonctionnalité**
- Vérifie si la chaîne est au format `IP1-IP2`.
- Valide que `IP1` et `IP2` sont des adresses IPv4 valides.
- Vérifie que `IP1` est inférieure à `IP2`.

#### **Exemple d'utilisation**

```go
type ExampleStruct struct {
	Range string `validate:"ipv4_range"` // Vérifie si la chaîne est une plage IPv4 valide
}
```

---

### **`TCPUDPPort`**

Le validateur `TCPUDPPort` vérifie si une valeur représente un port TCP ou UDP valide.

#### **Définition**

```go
var TCPUDPPort = &CustomValidator{
	Key: "tcp_udp_port",
	Func: func(fl validator.FieldLevel) bool {
		if fl.Field().IsZero() {
			return false
		}

		var port int
		switch fl.Field().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			port = int(fl.Field().Int())
		case reflect.String:
			i, err := strconv.Atoi(fl.Field().String())
			if err != nil {
				return false
			}
			port = i
		}

		if port <= 0 || port > 65535 {
			return false
		}

		return true
	},
}
```

#### **Fonctionnalité**
- Vérifie si la valeur est un entier ou une chaîne convertible en entier.
- Valide que la valeur est comprise entre `1` et `65535`.

#### **Exemple d'utilisation**

```go
type ExampleStruct struct {
	Port int `validate:"tcp_udp_port"` // Vérifie si la valeur est un port TCP/UDP valide
}
```

---

### **`TCPUDPPortRange`**

Le validateur `TCPUDPPortRange` vérifie si une chaîne de caractères représente une plage valide de ports TCP ou UDP.

#### **Définition**

```go
var TCPUDPPortRange = &CustomValidator{
	Key: "tcp_udp_port_range",
	Func: func(fl validator.FieldLevel) bool {
		p := strings.Split(fl.Field().String(), "-")
		if len(p) != 2 {
			return false
		}

		start, err := strconv.Atoi(p[0])
		if err != nil {
			return false
		}

		end, err := strconv.Atoi(p[1])
		if err != nil {
			return false
		}

		if start <= 0 || start > 65535 || end <= 0 || end > 65535 {
			return false
		}

		if start >= end {
			return false
		}

		return true
	},
}
```

#### **Fonctionnalité**
- Vérifie si la chaîne est au format `PORT1-PORT2`.
- Valide que `PORT1` et `PORT2` sont des entiers compris entre `1` et `65535`.
- Vérifie que `PORT1` est inférieur à `PORT2`.

#### **Exemple d'utilisation**

```go
type ExampleStruct struct {
	PortRange string `validate:"tcp_udp_port_range"` // Vérifie si la chaîne est une plage de ports valide
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
	IPv4Range string `validate:"ipv4_range"`
	Port      int    `validate:"tcp_udp_port"`
	PortRange string `validate:"tcp_udp_port_range"`
}

func main() {
	validate := validator.New()

	// Enregistrer les validateurs personnalisés
	validate.RegisterValidation(validators.IPV4Range.Key, validators.IPV4Range.Func)
	validate.RegisterValidation(validators.TCPUDPPort.Key, validators.TCPUDPPort.Func)
	validate.RegisterValidation(validators.TCPUDPPortRange.Key, validators.TCPUDPPortRange.Func)

	// Exemple de structure
	example := ExampleStruct{
		IPv4Range: "192.168.0.1-192.168.0.100",
		Port:      8080,
		PortRange: "1000-2000",
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

## **Licence**

Ce fichier est distribué sous la licence **Mozilla Public License 2.0**.  
Pour plus d'informations, consultez le fichier [LICENSE](https://www.mozilla.org/en-US/MPL/2.0/).

---

## **Résumé**

Le fichier network.go fournit trois validateurs personnalisés pour les adresses IPv4, les ports TCP/UDP et les plages de ports. Ces validateurs sont utiles pour garantir que les champs respectent des formats réseau spécifiques dans les structures Go.