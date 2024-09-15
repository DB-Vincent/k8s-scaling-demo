import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

export default async function Home() {
  let data = await fetch('http://localhost:8080/pods')
  let pods = await data.json()


  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">

      {pods.relatedPods.map((pod: string) => (
          <Card key={pod}>
            <CardHeader>
              <CardTitle>{pod}</CardTitle>
              <CardDescription>Some more information can go here</CardDescription>
            </CardHeader>
          </Card>
        ))}
      </main>
    </div>
  );
}
