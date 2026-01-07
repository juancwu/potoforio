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
					{TechName: "Next.js", TechURL: "https://nextjs.org/", TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-zinc-100 hover:text-gray-950"},
					{TechName: "Drizzle ORM", TechURL: "https://orm.drizzle.team/", TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-green-600"},
					{TechName: "Turso", TechURL: "https://turso.tech/", TechClasses: "transition bg-zinc-950 text-zinc-100 hover:bg-teal-600"},
				},
			},
			{
				Name:        "This Site",
				URL:         "https://juancwu.dev",
				Description: "My personal portolio site?",
				Repo:        "https://github.com/juancwu/potoforio",
				Classes:     "border-gray-600",
				TechList: []Tech{
					{TechName: "Golang", TechURL: "https://go.dev/", TechClasses: "bg-zinc-950 text-zinc-100"},
					{TechName: "HTMX", TechURL: "https://htmx.org/", TechClasses: "bg-zinc-950 text-zinc-100"},
					{TechName: "Hyperscript", TechURL: "https://hyperscript.org/", TechClasses: "bg-zinc-950 text-zinc-100"},
				},
			},
		},
	}
}
