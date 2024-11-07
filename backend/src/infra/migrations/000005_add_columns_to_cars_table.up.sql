ALTER TABLE cars
    ADD COLUMN price_for_initial_twelve_hours INTEGER DEFAULT 13000 NOT NULL,
    ADD COLUMN price_per_additional_six_hours INTEGER DEFAULT 6000 NOT NULL;

