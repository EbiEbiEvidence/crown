CREATE TABLE users
(
    user_id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR(256) UNIQUE NOT NULL,
    display_name VARCHAR(32) NOT NULL,
    token TEXT NOT NULL
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE games
(
    game_id SERIAL NOT NULL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(user_id),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    score INTEGER NOT NULL DEFAULT 0,
    emperor_imina VARCHAR(16) NOT NULL,
    emperor_gengo VARCHAR(16) NOT NULL
);

CREATE TABLE cards
(
    card_id SERIAL NOT NULL PRIMARY KEY,
    is_child BOOLEAN NOT NULL,
    scenario_id INTEGER,
    card_text VARCHAR(256) NOT NULL,
    card_image VARCHAR(16) NOT NULL
);

CREATE TABLE choices
(
    choice_id SERIAL NOT NULL PRIMARY KEY,
    next_card_id SERIAL REFERENCES cards(card_id),
    parent_card_id SERIAL REFERENCES cards(card_id),
    parent_card_direction VARCHAR(8) NOT NULL,
    choice_text VARCHAR(256) NOT NULL,
    church_modifier INTEGER NOT NULL,
    commoners_modifier INTEGER NOT NULL,
    merchants_modifier INTEGER NOT NULL,
    military_modifier INTEGER NOT NULL,
    bonus_modifier INTEGER NOT NULL,
    death_penalty BOOLEAN NOT NULL
);

CREATE TABLE game_choices
(
    game_choices_id SERIAL NOT NULL PRIMARY KEY,
    game_id SERIAL REFERENCES games(game_id),
    choice_id SERIAL REFERENCES choices(choice_id)
);
