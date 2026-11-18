const pengurus = [
  {
    name: "Nama Pengurus 1",
    role: "Ketua Umum",
    image: "/images/avatar-1.jpg",
  },
  {
    name: "Nama Pengurus 2",
    role: "Wakil Ketua",
    image: "/images/avatar-2.jpg",
  },
  {
    name: "Nama Pengurus 3",
    role: "Sekretaris",
    image: "/images/avatar-3.jpg",
  },
  { name: "Nama Pengurus 4", role: "Bendahara", image: "/images/avatar-4.jpg" },
];

export function PengurusSection() {
  return (
    <section id="pengurus" className="container pengurus-section">
      <div className="pengurus-heading-row">
        <div className="section-label">Pengurus</div>
        <h2 className="pengurus-heading">Kenali pengurus DOSCOM</h2>
      </div>
      <div className="pengurus-list">
        {pengurus.map((person) => (
          <article className="pengurus-person" key={person.name}>
            <img src={person.image} alt="" />
            <div>
              <h3>{person.name}</h3>
              <p>{person.role}</p>
            </div>
          </article>
        ))}
      </div>
    </section>
  );
}
