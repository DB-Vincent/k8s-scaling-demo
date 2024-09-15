import {
  Card,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import type { InferGetServerSidePropsType, GetServerSideProps } from 'next'

type Replicas = {
  currentPod: string
  relatedPods: string[]
}

export const getServerSideProps = (async () => {
  // Fetch data from external API
  const res = await fetch('http://localhost:8080/pods')
  const replicas: Replicas = await res.json()
  // Pass data to the page via props
  return { props: { replicas } }
}) satisfies GetServerSideProps<{ replicas: Replicas }>

export default function Home({
  replicas,
}: InferGetServerSidePropsType<typeof getServerSideProps>) {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        {replicas.currentPod}

          <Card>
            <CardHeader>
              <CardTitle>beep</CardTitle>
              <CardDescription>A specific endpoint</CardDescription>
            </CardHeader>
          </Card>
      </main>
    </div>
  );
}
