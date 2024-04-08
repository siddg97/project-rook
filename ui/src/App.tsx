import { Button, Card, CardBody, Tooltip } from '@nextui-org/react'

function App() {
  return (
    <>
      <div className="flex flex-wrap gap-4 items-center">
        <Tooltip content="I am a tooltip">
          <Button color="primary">
            Button
          </Button>
        </Tooltip>
        <Card>
          <CardBody>
            <p>This is a card</p>
          </CardBody>
        </Card>
      </div>
    </>
  )
}

export default App
