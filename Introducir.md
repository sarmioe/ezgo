# EasyGo es un programa que hace que Git y Mo CloudSpace sean fáciles de usar.

> La razón principal es que me volví loco con Git cuando lo usé para enviar MoMit, Me enojé por lo difícil que era usar Git, así que inventé esto
>
> Pero ahora, Mo Cloud Space sigue siendo un producto conceptual, por supuesto que puedes usar otros espacios en la nube.
>
> Veamos cómo usarlo. Primero debes descargar Git y Go y luego configurar las variables de entorno

| comando | Función |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| ezgo v | Mostrar versión |
| ezgo h | Imprimir ayuda predeterminada Inglés |
| ezgo h zh_CN | Imprimir ayuda de Chino |
| ezgo h zh_TW | Imprimir ayuda de Chino tradicional |
| ezgo h es | Imprimir ayuda de Español |
| ezgo update [versión] | Actualizar EasyGo |
| ezgo clone [URL] [ruta local] -branch--[nombre de la rama] -depth--[número] | Clonar repositorio desde la nube |
| ezgo sync [ruta local] [URL] | Ejecutar sincronización |
| ezgo sync auto [el tiempo predeterminado es segundo] | Sincronización automática |
| ezgo sync incremental | Sincronizar solo archivos de diferencia |
| ezgo config | Configurar EasyGo |
| ezgo env | Verificación automática del entorno |
| ezgo logs [nivel] | Salida de registros de ezgo |
| ezgo logs git | Salida de registros de git |
| ezgo logs go | Salida de registros de Go |
| ezgo push [confirmar] | Confirmar en repositorio remoto |
| ezgo pull [rama] | Extraer una rama |
| ezgo checkout [nombre de la rama] | Cambiar nombre de rama |
| ezgo conflict [vía] | Resolver conflictos locales y en la nube |
| ezgo go build [ARCH] [System] [Output file name] | Compilar archivos Go Compilar solo Go |
| ezgo go build all [Nombre del archivo de salida] | Compilar todas las versiones de arquitectura de todos los sistemas operativos Compilar solo Go |
| ezgo go init [nombre] | Inicializar proyecto Go |
| ezgo cs sync [nombre] | Sincronizar con el espacio en la nube Mo |
| ezgo cs config | Configurar la cuenta y la contraseña de la nube Mo |
| ezgo cs create [nombre] | Crear un espacio en la nube Mo |
| ezgo cs build [Entorno de idioma] [Nombre del archivo de salida] | Especificar un entorno de lenguaje y luego compilar. El entorno se puede ver en cs.clauded.modiznodz.llc |
| ezgo cs download [nombre] [path] | Descargar archivos desde la nube a un directorio local. Descarga cifrada por defecto |
| ezgo cs delete [nombre] | Eliminar proyectos en la nube. Borrar y sobrescribir automáticamente 7 veces por defecto |
| ezgo cs archive [nombre]                                     | Archivar el espacio en la nube de Mo en el almacenamiento en la nube de Mo o descargarlo al local |
| ezgo cs fork [URL] [nombre]                                  | Bifurcar un proyecto desde otro repositorio a Mo Cloud       |
| ezgo cs copy [nombre] [name2]                                | Copiar un CloudSpace a otro vacío                            |
| ezgo cs depoly [nombre del proyecto] [entorno]               | Implementar un proyecto en el alojamiento sin servidor de Mo Cloud |
