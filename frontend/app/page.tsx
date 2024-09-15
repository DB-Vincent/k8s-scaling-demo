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
    <div className="flex h-screen justify-center items-center">
      <main className="flex flex-col gap-4 items-center sm:items-start">
        {pods.relatedPods.map((pod: string) => (
          <Card key={pod} className={pods.currentPod === pod ? "border-stone-500 shadow-lg shadow-stone-500/20" : ""}>
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
