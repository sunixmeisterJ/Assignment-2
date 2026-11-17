# Composite Pattern in Go

Dokumentenstruktur aus der Vorlesung, ausschließlich mit der Go-Standardbibliothek.

## Starten und testen

Voraussetzung: Go 1.20 oder neuer. Im Projektverzeichnis:

```sh
go run .
go test ./...
```

Das Programm gibt `160` aus: Der Hauptordner enthält Textdokumente mit den Größen 10 und 50 sowie einen Unterordner mit einem Textdokument der Größe 100.

## Composite-Rollen

| Rolle | Go-Typ | Aufgabe |
| --- | --- | --- |
| Component | `Doc` | Gemeinsame Schnittstelle `Size() int` |
| Leaf | `TxtDoc` | Liefert seine feste Größe |
| Composite | `Folder` | Enthält `[]Doc`; `AddDoc` fügt Kinder hinzu; `Size` summiert rekursiv |

Blätter und Ordner werden über dasselbe Interface behandelt. Ein leerer Ordner hat Größe 0. Die Struktur wird als Baum ohne Zyklen und ohne nil-Kinder aufgebaut. Go verwendet hierfür Interfaces und Structs, keine Klassenvererbung.

Die Tests prüfen Blattgrößen, leere Ordner, mehrere Verschachtelungsebenen, Änderungen an Unterordnern und die Beispielausgabe. Der verwendete Prompt steht unverändert in `PROMPT.md`.
