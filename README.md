> [!WARNING]
> Yali is still in early development; **do not** depend on this project yet.
>
> The first public release will be tagged as `v1.0.0-beta`.

# Yali

Yali is a library implementing attribute based access control.

## Features

-   Set of [built-in attributes](docs/builtin-attributes.md) for subjects, actions, objects, and environments that aim to cover most use cases
-   Simple [predicate grammar](docs/predicate-grammar.md) to express policy rules and policy matchers

## Quickstart

Depend on Yali.

```bash
go get -u github.com/ccunni3/yali
```

Define policies.

```json

```

Initialize Yali.

```go

```

Enforce policies.

```go

```

## Support

Please [raise an issue](https://github.com/ccunni3/yali/issues/new) if you found a bug or you have questions about the project. The [user guide](docs/user-guide.md) contains more detailed, non-reference information about how to use Yali and how to implement it in your project.

All exported types are documented _passim_ as comments in the source code, conforming to the conventions set forth by the canonical blog post: [Go Doc Comments](https://go.dev/doc/comment).

## Contributing

Pull requests are welcome. Take a look for [open issues labeled "bug"](https://github.com/ccunni3/yali/labels/bug) and [milestones](https://github.com/ccunni3/yali/milestones) for more information about where the project is headed. Please initiate conversations about major changes by raising an issue.

If you're interested, read the [contribution guidelines](CONTRIBUTING.md) for this project.

## Acknowledgements

This project was inspired by [Casbin](https://casbin.org). I also thank the folks who design open specifications, especially those working on attribute based role control.

Yali was created by [Cole Cunningham](https://colecunningham.dev) in November 2023.

## License

This project is MIT licensed. See [LICENSE](LICENSE).

## Onomastic Explaination

I was inspired by the several examples of mythological creatures that protect entry or access to places while thinking of what to name the project. These mythical creatures often symbolize protection and the boundaries between different realms, serving as guardians of the physical and spiritual spaces they inhabit.

In Hindu mythology, the Yali is a mythical creature with the body of a lion or a tiger and the head of another animal, often an elephant. Yalis are frequently depicted as guardians of entrances to temples and other sacred places in South India.

![Sculpture of a Yali, a mythical creature with the body of a lion or a tiger and the head of another animal, often an elephant.](https://mapacademy.io/wp-content/uploads/2022/04/yali-vyala-2l.jpg)
