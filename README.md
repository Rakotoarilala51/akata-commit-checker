# AKATA Commit Checker - Documentation

## Description

**AKATA Commit Checker** est un outil CLI (Command Line Interface) développé en Go pour évaluer la qualité des messages de commit Git selon les conventions de commit d'AKATA GOAVANA. L'outil analyse les commits et attribue un score de qualité basé sur le respect des standards de formatage.

## Installation

### Prérequis
- Go 1.19+ installé sur votre système
- Git installé et configuré
- Un dépôt Git initialisé

### Compilation
```bash
# Cloner le projet
git clone https://github.com/Rakotoarilala51/akata-commit-checker
cd akata-commit-checker

# Compiler l'exécutable
go build -o check-commit .

# Rendre l'exécutable accessible globalement (optionnel)
sudo mv check-commit /usr/local/bin/
```

## Format de Commit Attendu

L'outil vérifie la conformité aux conventions de commit suivantes :

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types Valides
- `build` : Changements affectant le système de build
- `ci` : Changements dans la configuration CI/CD
- `docs` : Documentation uniquement
- `feat` : Nouvelle fonctionnalité
- `fix` : Correction de bug
- `perf` : Amélioration des performances
- `refactor` : Refactoring de code
- `style` : Changements de formatage/style
- `test` : Ajout ou modification de tests

### Exemples de Commits Valides

**Commit basique (Score: 3/5)**
```
feat: add user authentication
```

**Commit avec scope (Score: 4/5)**
```
feat(auth): add user authentication system
```

**Commit complet (Score: 5/5)**
```
feat(auth): add user authentication system

Implement JWT-based authentication with refresh tokens.
Add middleware for route protection and user session management.

Closes #123
```

## Utilisation

### Commandes Principales

#### 1. Analyser tous les commits
```bash
check-commit all
```
Analyse tous les commits dans toutes les branches du dépôt.

#### 2. Analyser une branche spécifique
```bash
check-commit branch <nom-branche>
```

**Exemples :**
```bash
check-commit branch main
check-commit branch develop
check-commit branch feature/user-auth
```

### Options Globales

#### Seuil de qualité (`--threshold` / `-t`)
Définit le score minimum requis pour considérer la qualité comme acceptable.

```bash
# Seuil par défaut (3/5)
check-commit all

# Seuil personnalisé (4/5)
check-commit all --threshold 4
check-commit all -t 4

# Seuil strict (5/5)
check-commit branch main --threshold 5
```

#### Mode verbeux (`--verbose` / `-v`)
Active l'affichage de détails supplémentaires lors de l'analyse.

```bash
check-commit all --verbose
check-commit branch main -v
check-commit all -t 4 -v
```

### Exemples d'Utilisation Complète

```bash
# Analyse simple de tous les commits
check-commit all

# Analyse stricte de la branche main
check-commit branch main --threshold 5

# Analyse détaillée avec seuil personnalisé
check-commit all --threshold 4 --verbose

# Analyse d'une branche de fonctionnalité
check-commit branch feature/payment-integration -t 3 -v
```

## Système de Scoring

L'outil attribue un score de 0 à 5 basé sur les critères suivants :

| Score | Statut | Critères |
|-------|--------|----------|
| **0/5** | ❌ **INVALID** | Format de commit invalide |
| **3/5** | ⚠️ **BASIC** | Type + Subject uniquement |
| **4/5** | ✅ **GOOD** | Type + Subject + (Scope OU Body OU Footer) |
| **5/5** | 🎉 **EXCELLENT** | Type + Subject + Scope + Body + Footer |

### Détail des Éléments

- **Type** *(obligatoire)* : Un des types valides listés ci-dessus
- **Subject** *(obligatoire)* : Description courte du changement
- **Scope** *(optionnel)* : Portée du changement (ex: `auth`, `api`, `ui`)
- **Body** *(optionnel)* : Description détaillée du changement
- **Footer** *(optionnel)* : Références aux issues, breaking changes, etc.

## Sortie et Codes de Retour

### Codes de Sortie
- **0** : ✅ Succès - Tous les commits respectent le seuil de qualité
- **1** : ❌ Échec - Un ou plusieurs commits ne respectent pas le seuil

### Utilisation dans CI/CD

```bash
#!/bin/bash
# Script de validation pour CI/CD

echo "🔍 Vérification de la qualité des commits..."

if check-commit branch main --threshold 4; then
    echo "✅ Qualité des commits validée"
    exit 0
else
    echo "❌ Qualité des commits insuffisante"
    exit 1
fi
```

### Exemples de Sortie

**Analyse Réussie :**
```
╔════════════════════════════════════════════════════════════╗
║                    COMMIT QUALITY ANALYSIS                 ║
╠════════════════════════════════════════════════════════════╣
║  Status: EXCELLENT COMMIT QUALITY                          ║
║  Subject: add user authentication system                   ║
║  Quality Score: 5/5                                        ║
╚════════════════════════════════════════════════════════════╝

╔══════════════════════════════════════════════════╗
║              COMMIT ANALYSIS RESULTS             ║
╠══════════════════════════════════════════════════╣
║ Total commits............: 0025                  ║
║ Valid commits............: 0023                  ║
║ Invalid commits..........: 0002                  ║
║ Average quality..........: 4.12/5                ║
║ Repository score.........: 4/5                   ║
║ Status: [ PASS ] Quality threshold met           ║
╚══════════════════════════════════════════════════╝

🎉 ANALYSE RÉUSSIE - Code de sortie: 0
```

## Dépannage

### Erreurs Courantes

**Erreur : "impossible d'ouvrir le dépôt Git"**
```bash
# Solution : Vérifier que vous êtes dans un dépôt Git
git status
```

**Erreur : "branche 'xxx' introuvable"**
```bash
# Solution : Vérifier les branches disponibles
git branch -a

# Utiliser le nom exact de la branche
check-commit branch feature/user-auth
```

**Format de commit invalide**
```bash
# Mauvais format
git commit -m "ajout authentification"

# Bon format
git commit -m "feat(auth): add user authentication system"
```

## Intégration avec Git Hooks

### Pre-commit Hook
Créer `.git/hooks/pre-commit` :

```bash
#!/bin/bash
# Vérifier le dernier commit avant push
if ! check-commit branch $(git branch --show-current) --threshold 3; then
    echo "❌ Commit rejeté : qualité insuffisante"
    exit 1
fi
```

### Pre-push Hook
Créer `.git/hooks/pre-push` :

```bash
#!/bin/bash
# Vérifier tous les commits avant push
if ! check-commit all --threshold 4; then
    echo "❌ Push rejeté : qualité globale insuffisante"
    exit 1
fi
```

## Support et Contribution

- **Repository** : https://github.com/Rakotoarilala51/akata-commit-checker
- **Issues** : Signaler les bugs ou demander des fonctionnalités
- **Contributions** : Les pull requests sont les bienvenues

---

*Documentation générée pour AKATA Commit Checker v1.0*
