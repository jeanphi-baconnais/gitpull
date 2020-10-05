# Git-pull  🧶

Simple script permettant de puller tous ses repositories git. 

Développé en Go,  🚧  en cours de première release.  🚧 

L'outil se base sur un répertoire, et détecte les repos Git. Pour chaque repo trouvé, il va récupérer les branches en cours et faire un _pull_.

## Installation 

### Mac OS : 

```
curl https://gitlab.com/jeanphi.baconnais/git-pull/-/raw/master/bin/mac/git-pull -o git-pull

chmod +x git-pull
```
### Linux : 


```
curl https://gitlab.com/jeanphi.baconnais/git-pull/-/raw/master/bin/linux/git-pull -o git-pull

chmod +x git-pull
```


## Utilisation 

Deux commandes sont diponibles :

- git-pull help : pour avoir une rapide explication de l'outil
- git-pull init : pour initialiser ou réinitialiser l'outil. Il s'agit simplement de saisir son répertoire contenant tous ses repositories git
