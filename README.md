# Curso de Go

Programa aplicaciones backend con Go. Aprende variables, funciones, goroutines y manejar concurrency para procesos eficientes. Desarrolla con GoMod y aprovecha paquetes externos. Implementa prácticas recomendadas desde métodos hasta channels.

## Clase 1: Características y ventajas del lenguaje de programación Go

Resumen

Go, también conocido como Golang, es un lenguaje de programación creado en 2007 por Robert Grismer, Bob Pike y Ken Thompson en Google, concebido como alternativa frente a la complejidad de lenguajes existentes como C++ y Java. Este lenguaje nació de la necesidad real de simplificar el desarrollo de software sin perjudicar el rendimiento y logró convertirse rápidamente en una opción muy valorada por empresas tecnológicas líderes.

¿Por qué elegir Go para desarrollar software?
Go ofrece características concretas que hacen que sea atractivo para empresas y desarrolladores. Específicamente:

Concurrencia sencilla: Permite ejecutar varias tareas simultáneamente de manera eficiente.
Compilación rápida: Reduce considerablemente el tiempo que se tarda en compilar un programa.
Sintaxis clara: Inspirada en C y Python, lo que facilita su aprendizaje y manejo.
Estas características se traducen en mejoras directas en la eficiencia de la infraestructura y reducción de costos operativos, algo confirmado por empresas que lo utilizan ampliamente.

¿Qué impacto tiene Go en empresas tecnológicas líderes?
Empresas destacadas de diversos sectores tecnológicos han adoptado Go, confirmando su eficiencia operativa mediante datos concretos. Por ejemplo:

MercadoLibre: Redujo un 88% el uso de servidores y a la mitad el consumo de CPU al migrar sus sistemas.
Uber, Netflix, PayPal y Twitch: Estas compañías han elegido Go debido a la necesidad de contar con servicios rápidos, escalables y confiables.
Estas implementaciones muestran claramente la utilidad práctica de Go en resolver problemas complejos manteniendo un bajo costo operativo.

¿Qué ofrece la comunidad de Go?
Al aprender Go, los desarrolladores no solo acceden a otro lenguaje, sino que ingresan en una comunidad activa, colaborativa y con un fuerte enfoque en la mejora constante:

Un ecosistema creciente lleno de bibliotecas y herramientas útiles.
Una comunidad colaborativa, vigente y en expansión, que facilita el aprendizaje continuo y la mejora de habilidades.
Una filosofía claramente definida que apuesta por la simplicidad y eficiencia tecnológica, simbolizada por la mascota oficial del lenguaje, el Gopher.
Go es, por lo tanto, más que solo código; representa un estilo de desarrollo orientado al máximo rendimiento con el uso mínimo de recursos.

## Clase 2: Instalación de Go en Ubuntu con WSL y Visual Studio Code

Resumen

¿Quieres empezar a programar en Go rápidamente? La instalación ahora es más sencilla gracias a WSL y el administrador de paquetes Snap en Ubuntu, facilitando que en solo un comando tengas todo listo para trabajar. Además, complementado con Visual Studio Code y la extensión oficial de Go, puedes mejorar enormemente tu experiencia de desarrollo.

¿Cómo instalar Go en Ubuntu usando Snap y WSL?
Instalar Go en tu máquina a través de WSL con Ubuntu es simple. Solo abre tu terminal y ejecuta el siguiente comando:

sudo snap install go
Luego, confirma la instalación ejecutando:

go version
Gracias a Snap, Go se actualizará automáticamente, lo que simplifica aún más el mantenimiento.

¿Qué ventajas ofrece utilizar Visual Studio Code para desarrollar en Go?
Visual Studio Code junto con su extensión oficial para Go, creada por el equipo de Google, ofrece funcionalidades esenciales como:

Mejor soporte para IntelliSense.
Eficaz marcado de sintaxis.
Indicaciones útiles para importar y usar paquetes necesarios.
Solo escribe "Go" en el buscador de extensiones e instala la proporcionada por Google para contar con estas facilidades.

¿Cuáles son los principales comandos para trabajar con Go?
Aprender a usar los comandos básicos es fundamental para trabajar en Go:

go run: Ejecuta directamente tu código sin generar archivos adicionales, ideal para entornos locales.
go build: Crea un archivo compilado que puedes ejecutar para pruebas más detalladas o para desplegar la aplicación en entornos remotos.
Ejemplo práctico para ejecutar un programa:

go run main.go
Para generar archivos ejecutables utiliza:

go build main.go
./main
Los comandos go run y go build serán de uso frecuente según tus necesidades específicas cuando desarrollas aplicaciones en Go.

Recuerda que a medida que avanzas, podrás profundizar aún más en cómo aprovechar al máximo todo lo que Go tiene para ofrecer. ¿Qué te ha parecido trabajar con Go y Visual Studio Code hasta ahora? ¡Comparte tu experiencia!

## Clase 3: Gestión de paquetes e imports en Go

Resumen

Go se destaca por su rapidez, eficiencia y tamaño reducido, características que derivan en particularidades clave que todo programador debería conocer. La utilización eficiente de paquetes y la gestión estricta de imports son elementos esenciales que diferencian a Go de lenguajes como Python, Node o Java.

¿Qué sucede específicamente con los paquetes en Go?
A diferencia de otros lenguajes como Python que cuentan con sistemas como pip, Node con npm o .Net con nugget, Go no utiliza un manejador de paquetes externo. En Go, los paquetes se importan directamente según su utilidad específica, asegurando así la optimización del rendimiento.

Al importar elementos como FMT y OS, Go automáticamente verifica qué paquetes se están utilizando realmente. Si se importa un paquete y no se utiliza, Go lo señala explícitamente mediante avisos, lo que previene errores y optimiza el tamaño de compilación.

¿Cómo gestionar adecuadamente los paquetes?
Para importar paquetes en Go correctamente:

Usa la sintaxis apropiada con la palabra clave import.
Evita cargar paquetes que no se usarán realmente.
Revisa constantemente los mensajes del IDE donde Go indica paquetes sin utilizar.
Por ejemplo:

```Go
package main

import (
    "fmt"
    "os"
)

func main() {
    envVar := os.Getenv("HOME")
    if envVar == "" {
        fmt.Println("La variable HOME no está configurada")
    } else {
        fmt.Println("HOME sí está configurada y es:", envVar)
    }
}
```

¿Cuál es la utilidad práctica del uso estricto de paquetes?
La optimización de paquetes tiene efectos inmediatos y visibles:

Menor tamaño y mayor velocidad en la aplicación final.
Mejora notoria en rendimiento, tanto en compilación como en ejecución.
Código más limpio, claro y conciso.
¿Qué sucede si importamos paquetes innecesarios?
Go no permite compilar códigos con paquetes extras que no estén siendo utilizados directamente, mostrando un aviso específico en el IDE y evitando que la compilación se concrete, promoviendo así un rendimiento óptimo.

¿Cómo crear archivos u obtener variables ambientales utilizando paquetes?
Usando el paquete os se puede acceder a variables ambientales fácilmente, además de gestionar archivos nuevos directamente desde el código escrito.

Ejemplo para generar un archivo desde Go:

package main

```Go
import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("ejemplo.txt")
    if err != nil {
        fmt.Println("Error en creación del archivo:", err)
        return
    }
    defer file.Close()
    fmt.Println("Archivo creado exitosamente")
}
```

Esta gestión directa refuerza un control estricto sobre las dependencias utilizadas dentro del código, garantizando que solamente se incluya aquello que se necesita expresamente.

Cuéntame, ¿habías observado ya esta gestión de paquetes en tus proyectos con Go?
