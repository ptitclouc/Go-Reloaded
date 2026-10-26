# Go-Reloaded

tu dois avoir pour lancer mon programme vscode, git et golang.

lien à cloner (avec git clone):
https://github.com/ptitclouc/Go-Reloaded.git

Mon programme corrige les erreurs dans un texte:
- 01.fichier.go comporte les fonctions entrée et sortie qui lis et écris un .txt
- 02_découpage.go comporte la fonction decoupage qui coupe le texte en mots à chaque espace
- 03_hex_bin.go comporte la fonction hexbin qui convertit le mot d'avant en nombre normal quand on met (hex) ou (bin)
- 04_up_low_cap.go comporte les fonctions Up, Low et Cap qui mettent le mot d'avant en majuscules, en minuscules ou avec une majuscule au début
- 05_ponctuation.go comporte la fonction ponctuation qui colle . , ! ? : ; au mot d'avant
- 06_apostrophes.go comporte la fonction apostrophes qui colle les apostrophes aux mots à l'intérieur 
- 07_a_an.go comporte la fonction atooan qui change "a" en "an" devant une voyelle (pas encore fini)
- 08_tests_test.go comporte les tests qui vérifient que les fonctions marchent (à lancer avec go test)
- capitalize.go comporte la fonction Capitalize qui met la première lettre d'un mot en majuscule et le reste en minuscule, utilisée par la func Cap
- main.go appelle toutes les fonctions dans l'ordre et écrit le texte corrigé
- go.mod` définit le module Go du projet
- sample.txt est le texte de départ pour tester
- result.txt est le texte corrigé écrit par le programme

Alexandre DEBRABANT georges armand
