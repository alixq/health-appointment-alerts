# Description

This is a small CLI script I wrote to make getting vaccine appointments easier in France, where most of the bookings are centralized on a single platform.
For this reason, the rest of this document will be in French.

# Compatibilité

Pour le moment, ça ne marche que sur les Mac.

# Installation

Le mieux est d'[installer go](https://golang.org/) sur votre machine, après quoi vous pouvez ouvrir le terminal et éxécuter:

```
go install "github.com/alixq/rdv-sante@latest"
rdv-sante
```

Pour le moment je ne gère pas très bien le versioning, si besoin d'installer la dernière version:

```
rm ~/go/bin/rdv-sante
go install "github.com/alixq/rdv-sante@latest"
```

# Guide

Une fois l'application lancée, vous devrez rentrer un URL de recherche doctolib. Une fenêtre s'ouvrira quand vous ferez entrée,
vous tapez votre recherche sur le site et une fois que vous avez accès aux résultats copiez-collez l'URL dans la fenêtre.

Ensuite l'application chargera les centres de santé des prochaines pages de résultats. Vous pourrez les sélectionner un par un en
fonction de ce qui vous arrange.

Une fois sélectionné, l'application va solliciter régulièrement Doctolib pour vérifier les rendez-vous.

Lors d'un lancement subséquent de l'application, votre recherche aura été enregistrée et vous pourrez reprendre votre config précédente,
ou en refaire une nouvelle.

Pour le moment, la config est plutôt basée sur une utilisation pour trouver un RDV vaccin, même si théoriquement c'est utilisable
pour d'autres types de recherches doctolib. Les requêtes sont relancées très régulièrement pour s'adapter à la demande énorme,
et pour être sur de ne pas rater l'event votre Mac vous notifiera de la disponibilité d'un rendez-vous, et la fenêtre pour le réserver
s'ouvrira. Aussi, seuls les rendez-vous du jour même ou du lendemain sont notifiés pour le moment car c'est la guideline officielle pour
les 18-49 ans sans morbidité. Cela pourra être customisé à l'avenir pour être compatible avec une population plus diverse.

C'est une UX absolument pourrie, mais pour le usecase vaccins, c'est nécessaire. Une version viendra pour laquelle on pourra customiser
la manière de notifier, et la fréquence du ping.

Autre chose: **souvent, une fenêtre va s'ouvrir et il n'y aura plus rien**. C'est comme ça, sûrement parce-que quelqu'un est
plus malin ou plus réactif. Prenez votre mal en patience, normalement vous devriez en trouver un.

Lorsque une fenêtre s'ouvre et que la voix se déclenche, le script est arrêté. Si vous n'avez pas trouvé votre bonheur ou que vous
souhaitez réserver d'autres rendez-vous, il faut relancer. Autrement vous auriez une fenêtre qui s'ouvre par seconde et un son très
désagréable à l'oreille.
