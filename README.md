# hexagonal-architecture
Exemplo de aplicação utilizando os conceitos de arquitetura hexagonal na linguagem Go

[![CI Golang](https://github.com/jonathanmdr/hexagonal-architecture/actions/workflows/ci-go.yml/badge.svg?branch=master)](https://github.com/jonathanmdr/hexagonal-architecture/actions/workflows/ci-go.yml)

---

A arquitetura hexagonal consiste no modelo de portas e adaptadores, onde separamos completamente nosso domínio de negócio de componentes externos, permitindo um alto nível de desacoplamento.

Em um primeiro momento a percepção pode parecer a de implementar muito código para pouco resultado, porém, o ganho real vem com o tempo, onde se torna perceptível a velocidade que ganhamos no desenvolvimento, torna a manutenibilidade e evolução da aplicação mais simples e além disso, um outro fator interessante é que com este padrão arquitetural conseguimos plugar e desplugar qualquer componente externo ao nosso domínio com o mínimo de esforço e sem grandes preocupações.
