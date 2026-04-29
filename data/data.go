package data

type Tech struct {
	TechName    string
	TechURL     string
	TechClasses string
}

type Project struct {
	Name        string
	URL         string
	Description string
	TechList    []Tech
	Repo        string
	Classes     string
}

type Page struct {
	Projects []Project
}

func GetPageData() Page {
	return Page{
		Projects: []Project{
			{
				Name:        "Splinter",
				URL:         "",
				Description: "Simple and minimal structured logger in Go with file rotation.",
				Repo:        "https://git.juancwu.dev/juancwu/splinter",
				Classes:     "border-yellow-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Errx",
				URL:         "",
				Description: "Go error library that builds readable breadcrumb chains as errors propagate.",
				Repo:        "https://git.juancwu.dev/juancwu/errx",
				Classes:     "border-pink-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Cubby",
				URL:         "",
				Description: "Tiny in-memory key-value cache shared between processes via a Unix socket.",
				Repo:        "https://git.juancwu.dev/juancwu/cubby",
				Classes:     "border-amber-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Ficha",
				URL:         "",
				Description: "Go library for opaque tokens with embedded permissions.",
				Repo:        "https://git.juancwu.dev/juancwu/ficha",
				Classes:     "border-rose-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Lightmux",
				URL:         "",
				Description: "Small wrapper around Go's net/http ServeMux with groups and middleware.",
				Repo:        "https://git.juancwu.dev/juancwu/lightmux",
				Classes:     "border-emerald-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Ccretdns",
				URL:         "",
				Description: "DNS tool to resolve private domain names and public domains using upstream servers.",
				Repo:        "https://git.juancwu.dev/juancwu/ccretdns",
				Classes:     "border-red-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Gosh",
				URL:         "",
				Description: "Context-aware, just-in-time SSH agent wrapper.",
				Repo:        "https://git.juancwu.dev/juancwu/gosh",
				Classes:     "border-green-400",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Budgit",
				URL:         "",
				Description: "Simple expense tracker with real-time grocery list.",
				Repo:        "https://git.juancwu.dev/juancwu/budgit",
				Classes:     "border-blue-300",
				TechList: []Tech{
					{TechName: "Go", TechURL: "https://go.dev/", TechClasses: ""},
					{TechName: "PostgreSQL", TechURL: "https://postgresql.org", TechClasses: ""},
					{TechName: "TailwindCSS", TechURL: "https://tailwindcss.com/", TechClasses: ""},
				},
			},
			{
				Name:        "HawkHacks Website",
				URL:         "https://hawkhacks.ca",
				Description: "HawkHacks is a 36 hour in-person hackathon hosted at Wilfrid Laurier University.",
				Repo:        "https://github.com/LaurierHawkHacks/Landing",
				Classes:     "border-orange-300",
				TechList: []Tech{
					{TechName: "React", TechURL: "https://react.dev/", TechClasses: ""},
					{TechName: "Firebase", TechURL: "https://firebase.google.com/", TechClasses: ""},
					{TechName: "TailwindCSS", TechURL: "https://tailwindcss.com/", TechClasses: ""},
				},
			},
			{
				Name:        "Konfer Website",
				URL:         "https://konfer.ca",
				Description: "The official landing site for Konfer Inc.",
				Repo:        "https://github.com/KonferCA/Konfer.ca",
				Classes:     "border-indigo-300",
				TechList: []Tech{
					{TechName: "React", TechURL: "https://react.dev/", TechClasses: ""},
					{TechName: "Firebase", TechURL: "https://firebase.google.com/", TechClasses: ""},
					{TechName: "TailwindCSS", TechURL: "https://tailwindcss.com/", TechClasses: ""},
				},
			},
			{
				Name:        "LCS Website",
				URL:         "https://lauriercs.ca",
				Description: "The official Computer Science club at Wilfrid Laurier University website.",
				Repo:        "https://github.com/LaurierCS/Website",
				Classes:     "border-cyan-400",
				TechList: []Tech{
					{TechName: "React", TechURL: "https://react.dev/", TechClasses: ""},
					{TechName: "Firebase", TechURL: "https://firebase.google.com/", TechClasses: ""},
					{TechName: "Mantine UI", TechURL: "https://mantine.dev/", TechClasses: ""},
				},
			},
			{
				Name:        "Shoto",
				URL:         "",
				Description: "Simple and effective URL shortoner.",
				Repo:        "https://github.com/juancwu/shoto",
				Classes:     "border-indigo-500",
				TechList: []Tech{
					{TechName: "Next.js", TechURL: "https://nextjs.org/", TechClasses: ""},
					{TechName: "Drizzle ORM", TechURL: "https://orm.drizzle.team/", TechClasses: ""},
					{TechName: "Turso", TechURL: "https://turso.tech/", TechClasses: ""},
				},
			},
			{
				Name:        "This Site",
				URL:         "https://juancwu.dev",
				Description: "My personal portolio site?",
				Repo:        "https://github.com/juancwu/potoforio",
				Classes:     "border-gray-600",
				TechList: []Tech{
					{TechName: "Golang", TechURL: "https://go.dev/", TechClasses: ""},
					{TechName: "HTMX", TechURL: "https://htmx.org/", TechClasses: ""},
					{TechName: "Hyperscript", TechURL: "https://hyperscript.org/", TechClasses: ""},
				},
			},
		},
	}
}
