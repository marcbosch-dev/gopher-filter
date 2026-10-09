# gopher-filter

Aquí tienes el enunciado de requerimientos (**Subject**) enfocado exclusivamente en el comportamiento esperado de la aplicación:

---

### Proyecto: Procesador y Transformador de Datos en Lote

#### **Objetivo General**

Construir una herramienta por línea de comandos que lea un archivo de datos con múltiples registros (por ejemplo, transacciones financieras o datos de usuarios), valide su contenido, aplique una serie de transformaciones y filtros configurables, y entregue un reporte final junto con los datos procesados.

---

#### **Requerimientos Funcionales (Comportamiento Esperado)**

1. **Lectura e Ingesta de Datos**
* El programa debe recibir la ruta de un archivo de origen (CSV o JSON) que contenga una colección de registros con al menos 4 atributos de distintos tipos (ej. *ID, Nombre, Monto, Estado*).
* Debe verificar que el archivo exista y se pueda abrir antes de procesarlo.


2. **Criterios de Filtrado Dinámicos**
* El sistema debe permitir aplicar uno o varios criterios de selección sobre la lista de registros de manera acumulativa.
* *Ejemplos de criterios:* Filtrar por un valor mínimo/máximo en campos numéricos, filtrar por coincidencia exacta o parcial de cadenas de texto, o filtrar por estado activo/inactivo.
* Debe ser posible combinar múltiples filtros en una misma ejecución sin modificar el núcleo del motor.


3. **Pipeline de Transformación**
* El programa debe aplicar modificaciones a los registros que hayan superado los filtros.
* *Ejemplos de transformaciones:* Normalizar textos (ej. convertir emails a minúsculas), calcular nuevos valores sobre campos numéricos (ej. aplicar un recargo o impuesto), o enmascarar datos sensibles.
* Cada regla de transformación debe poder fallar individualmente si los datos no cumplen con las reglas de negocio.


4. **Manejo de Errores y Tolerancia a Fallos**
* Si un registro individual contiene un dato corrupto o inválido durante el proceso, el programa no debe detenerse; debe omitir o marcar ese registro, registrar el motivo del error y continuar procesando el resto.
* Si ocurre un error fatal al intentar abrir el archivo o procesar la entrada, el programa debe abortar de manera limpia mostrando un mensaje explicativo.


5. **Telemetría y Control de Ejecución**
* Al iniciar y finalizar cada etapa clave (lectura, filtrado, transformación y reporte), la aplicación debe imprimir en consola el tiempo exacto transcurrido en esa fase.
* Debe asegurarse de que los recursos del sistema (como descriptores de archivo abierto) se liberen siempre de forma garantizada, incluso si la ejecución se interrumpe por un fallo durante el proceso.


6. **Salida y Reporte Final**
* El programa debe mostrar en pantalla (o guardar en un archivo nuevo) la lista resultante de registros filtrados y transformados.
* Debe imprimir un resumen estadístico que incluya:
* Total de registros leídos.
* Total de registros que superaron los filtros.
* Total de registros transformados con éxito.
* Total de registros descartados o con errores.