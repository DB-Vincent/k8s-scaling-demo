import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

interface Replica {
  name: string;
  current: boolean;
  nodeName: string;
  status: string;
  startTime: string;
}

export default async function Home() {
  let data = await fetch('http://localhost:8080/pods')
  let pods = await data.json()

  return (
    <div className="flex h-screen justify-center items-center">
      <main className="flex flex-col gap-4 items-center sm:items-start">
        {pods.replicas.map((pod: Replica) => (
          <Card key={pod.name} className={pod.current ? "border-stone-500 shadow-lg shadow-stone-500/20" : ""}>
            <CardHeader>
              <CardTitle>{pod.name}</CardTitle>
              <CardDescription>{pod.status}</CardDescription>
            </CardHeader>
            <CardContent>
              <p>{pod.nodeName}</p>
              <p>{pod.startTime}</p>
            </CardContent>
          </Card>
        ))}
      </main>
    </div>
  );
}
