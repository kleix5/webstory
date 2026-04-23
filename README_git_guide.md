# Useful Git Guide

Практичная шпаргалка по Git на основе нашего чата и стартового файла `GitCheatSheet.txt`.

---

## Содержание

1. [Создание и инициализация](#создание-и-инициализация)
2. [Создание SSH-соединения](#создание-ssh-соединения)
3. [Работа с удалённым репозиторием](#работа-с-удалённым-репозиторием)
4. [Как подключиться к чужому удалённому репозиторию](#как-подключиться-к-чужому-удалённому-репозиторию)
5. [Управление проектом: PR, merge, diff, rebase](#управление-проектом-pr-merge-diff-rebase)
6. [Git Flow под твой проект (journal / hotel)](#git-flow-под-твой-проект-journal--hotel)
7. [Полезные мини-шпаргалки](#полезные-мини-шпаргалки)

---

## Создание и инициализация

### 1) Инициализация репозитория

```bash
git init -b main
```

Если репозиторий уже инициализирован:

```bash
git status
```

### 2) Первый коммит

Важно: до первого коммита многие операции с ветками либо бессмысленны, либо работают не так, как ожидает новичок.

```bash
git add .
git commit -m "Initial commit"
```

### 3) Создание рабочей ветки

Для твоего сценария:

```bash
git switch -c reassemble/journal-db-baseline
```

Смысл названия:
- `reassemble/` - префикс твоего типа работ
- `journal-db-baseline` - базовая сборка структуры БД проекта journal

### 4) Переключение между ветками

```bash
git branch
git switch main
git switch reassemble/journal-db-baseline
```

Быстро вернуться на предыдущую ветку:

```bash
git switch -
```

### 5) Что важно помнить про ветки

- Ветки независимы до merge.
- Удаление файла в рабочей ветке не удаляет его из `main`, пока ты не сделал merge.
- `main` лучше держать как стабильную ветку.
- Работать лучше не в `main`, а в отдельных ветках.

---

## Создание SSH-соединения

Тебе удобнее работать из Git Bash, поэтому здесь команды даны под Git Bash.

### 1) Создать SSH-ключ

```bash
ssh-keygen -t ed25519 -C "your_email@example.com"
```

На вопросы:
- путь сохранения - просто Enter
- passphrase - Enter для простого варианта или пароль для более защищённого

### 2) Проверить, что ключи появились

```bash
ls ~/.ssh
```

Обычно будут файлы:
- `id_ed25519`
- `id_ed25519.pub`

### 3) Показать публичный ключ

```bash
cat ~/.ssh/id_ed25519.pub
```

Копировать нужно всю строку целиком:
- `ssh-ed25519`
- длинный ключ
- комментарий в конце (обычно email)

### 4) Добавить ключ в GitHub

На GitHub:
- Settings
- SSH and GPG keys
- New SSH key
- вставить скопированную строку

### 5) Проверить соединение

```bash
ssh -T git@github.com
```

При первом подключении GitHub спросит подтверждение:

```text
yes
```

Успешный ответ будет примерно таким:

```text
Hi <username>! You've successfully authenticated, but GitHub does not provide shell access.
```

### 6) Главное понимание

SSH-ключ привязан не к одному проекту, а к твоему аккаунту.

То есть один ключ может использоваться для всех репозиториев, к которым у тебя есть доступ.

---

## Работа с удалённым репозиторием

### 1) Добавить удалённый репозиторий

SSH-вариант:

```bash
git remote add origin git@github.com:USERNAME/REPOSITORY.git
```

Проверка:

```bash
git remote -v
```

### 2) Что такое `origin`

`origin` - это просто имя удалённого репозитория.

Обычно используют именно `origin`, но это всего лишь alias.

### 3) Первый push ветки `main`

```bash
git push -u origin main
```

### 4) Push рабочей ветки

```bash
git switch reassemble/journal-db-baseline
git push -u origin reassemble/journal-db-baseline
```

### 5) Что означает `-u`

`-u` = `--set-upstream`

Это означает: Git запоминает связь между локальной веткой и удалённой.

После первого такого push можно чаще писать просто:

```bash
git push
git pull
```

### 6) Разница между этими командами

```bash
git push origin reassemble/journal-db-baseline
git push -u origin reassemble/journal-db-baseline
```

Разница такая:
- первая команда просто пушит
- вторая пушит и запоминает upstream-связь

---

## Как подключиться к чужому удалённому репозиторию

Это уже сценарий командной работы.

### Вариант 1) Ты участник репозитория

Если тебя добавили в проект, ты можешь:

```bash
git clone git@github.com:OWNER/REPO.git
cd REPO
git branch
git switch -c feature/my-task
```

Дальше работаешь в своей ветке:

```bash
git add .
git commit -m "Implement my task"
git push -u origin feature/my-task
```

Потом открываешь PR.

### Вариант 2) У тебя нет прав на запись

Тогда обычный путь такой:
1. сделать Fork на GitHub
2. клонировать уже свой форк
3. работать в своей ветке
4. пушить в свой fork
5. открыть PR из fork в оригинальный репозиторий

### Полезно знать про remotes в команде

Очень часто встречается такая схема:
- `origin` - твой fork
- `upstream` - оригинальный репозиторий команды

Добавление `upstream`:

```bash
git remote add upstream git@github.com:TEAM/PROJECT.git
git remote -v
```

Обновить локальный `main` из командного репозитория:

```bash
git fetch upstream
git switch main
git merge upstream/main
```

Когда освоишься лучше, можно делать и через rebase.

---

## Управление проектом: PR, merge, diff, rebase

### 1) Что такое PR

PR (Pull Request) - это предложение влить изменения из одной ветки в другую.

В твоём случае:

```text
reassemble/journal-db-baseline -> main
```

То есть ты говоришь:

> Я поработал в своей ветке. Посмотрите изменения и добавьте их в `main`.

### 2) Как выглядит рабочий цикл

#### Работа в своей ветке

```bash
git switch reassemble/journal-db-baseline
git add .
git commit -m "Add DB baseline"
git push
```

#### Создание PR на GitHub

На GitHub:
- base: `main`
- compare: `reassemble/journal-db-baseline`

#### После проверки

На GitHub нажимается кнопка Merge pull request.

### 3) Локальный merge

Если делаешь не через GitHub, а локально:

```bash
git switch main
git merge reassemble/journal-db-baseline
```

Важно:
- merge делается в той ветке, в которой ты сейчас находишься
- если ты стоишь в `main`, изменения попадут в `main`

### 4) `git diff`

Показать разницу между ветками:

```bash
git diff main..reassemble/journal-db-baseline
```

Показать незакоммиченные изменения:

```bash
git diff
```

Показать staged-изменения:

```bash
git diff --staged
```

### 5) `git restore --staged`

Если случайно добавил файл в staging:

```bash
git restore --staged file.txt
```

Это не удаляет изменения из файла. Команда только убирает файл из staging area.

### 6) Что такое `git rebase`

`rebase` - это способ "пересадить" свои коммиты на более свежую базу.

Пример:

У тебя есть ветка:

```text
main:      A---B---C
feature:        \---D---E
```

Потом в `main` появились новые коммиты, и ты хочешь, чтобы твоя ветка выглядела так, будто ты начал работу от самой свежей версии `main`:

```text
main:      A---B---C---F---G
feature:                \---D'---E'
```

Это и делает rebase.

### 7) Зачем нужен rebase

Он нужен, чтобы:
- подтянуть свежий `main` в свою ветку
- получить более линейную и чистую историю
- уменьшить хаос перед PR

### 8) Базовый rebase-сценарий

Когда ты работаешь в своей ветке и хочешь подтянуть свежий `main`:

```bash
git fetch origin
git switch reassemble/journal-db-baseline
git rebase origin/main
```

### 9) Если во время rebase возник конфликт

Git остановится. Тогда сценарий такой:

```bash
# исправить конфликт в файлах
git add .
git rebase --continue
```

Если хочешь отменить весь rebase:

```bash
git rebase --abort
```

### 10) Когда rebase особенно полезен

Полезно:
- в своей личной ветке до merge
- перед открытием PR
- чтобы обновить ветку после изменений в `main`

Опасно:
- делать rebase чужой общей ветки, где уже работают другие люди
- переписывать историю ветки, которую уже активно используют в команде

### 11) Простое правило

- `merge` - безопаснее и проще понять
- `rebase` - чище история, но требует внимательности

Практичный совет:
- сначала хорошо почувствуй `merge`
- потом используй `rebase` в личных рабочих ветках

---

## Git Flow под твой проект (journal / hotel)

Тебе пока не нужен тяжёлый классический Git Flow с `develop`, `release`, `hotfix` и кучей параллельных правил.

Для твоих проектов лучше использовать облегчённый и очень понятный flow.

### Вариант для одного разработчика или спокойной небольшой команды

```text
main
├── reassemble/journal-db-baseline
├── feature/booking-form
├── feature/admin-panel
├── bugfix/login-error
└── docs/git-guide
```

### Роли веток

#### `main`
Стабильная ветка.

Там должен лежать рабочий код.

#### `reassemble/*`
Для пересборки, крупных переустройств, архитектурных изменений.

Хорошо подходит для:
- новой структуры базы данных
- изменения каркаса проекта
- пересборки логики бронирования

#### `feature/*`
Для новых функций.

Примеры:
- `feature/guest-checkin-form`
- `feature/booking-calendar`
- `feature/payment-page`

#### `bugfix/*`
Для обычных исправлений.

Примеры:
- `bugfix/price-rounding`
- `bugfix/login-validation`

#### `docs/*`
Для документации и обучающих материалов.

Примеры:
- `docs/api-notes`
- `docs/git-guide`

### Рекомендуемый flow для тебя

#### Для проекта journal

```text
main
├── reassemble/journal-db-baseline
├── feature/student-profile
├── feature/teacher-dashboard
├── feature/attendance-table
└── bugfix/grade-save-error
```

#### Для проекта hotel

```text
main
├── reassemble/booking-core
├── feature/apartment-gallery
├── feature/corporate-booking
├── feature/admin-availability
└── bugfix/date-overlap
```

### Практика работы

1. В `main` не работаешь напрямую.
2. На каждую задачу создаёшь отдельную ветку.
3. Делаешь маленькие понятные коммиты.
4. Перед merge смотришь diff.
5. Лучше вливать изменения через PR.
6. Если `main` ушёл вперёд - подтягиваешь его в свою ветку через `merge` или `rebase`.

### Мини-шаблон на каждую задачу

```bash
git switch main
git pull
git switch -c feature/my-task

# работа
git add .
git commit -m "Implement my task"
git push -u origin feature/my-task
```

Потом:
- открыть PR
- проверить Files changed
- сделать merge
- удалить ветку

---

## Полезные мини-шпаргалки

### 1) Базовый сценарий с нуля

```bash
git init -b main
git add .
git commit -m "Initial commit"
git switch -c reassemble/journal-db-baseline
```

### 2) Отправка на GitHub

```bash
git remote add origin git@github.com:USERNAME/REPOSITORY.git
git push -u origin main
git push -u origin reassemble/journal-db-baseline
```

### 3) Merge своей ветки в main локально

```bash
git switch main
git merge reassemble/journal-db-baseline
```

### 4) Обновить свою ветку от свежего main через rebase

```bash
git fetch origin
git switch reassemble/journal-db-baseline
git rebase origin/main
```

### 5) Если случайно добавил лишнее в staging

```bash
git restore --staged file.txt
```

### 6) Полезные команды для проверки

```bash
git status
git branch
git log --oneline --decorate --graph -n 10
git remote -v
git diff
git diff --staged
```

---

## Главное, что стоит запомнить

- `git add` не делает commit.
- `git commit` фиксирует staged-изменения.
- ветки независимы до merge.
- PR - это запрос на вливание одной ветки в другую.
- `origin` - имя удалённого репозитория.
- `-u` при push запоминает upstream.
- SSH-ключ принадлежит твоему аккаунту, а не одному проекту.
- `merge` проще и безопаснее.
- `rebase` делает историю чище, но требует внимательности.

---

## Рекомендуемый личный стиль работы

Для твоего уровня и твоих проектов я бы советовал такой стиль:

1. `main` держать чистым.
2. На каждую задачу делать отдельную ветку.
3. Использовать `reassemble/*` для крупных переделок.
4. Использовать `feature/*` для новых функций.
5. Делать PR даже если работаешь один - ради самопроверки.
6. Перед merge смотреть diff и список удаляемых файлов.
7. Rebase применять в своих личных ветках, когда уже уверенно чувствуешь историю Git.

