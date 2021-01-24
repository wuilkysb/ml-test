CREATE TABLE IF NOT EXISTS public.mutants
(
    id              serial          NOT NULL,
    dna             text[]          NOT NULL,
    is_mutant       boolean         NOT NULL,
    PRIMARY KEY (id)
);