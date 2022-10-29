import logging
import datetime
from telegram import Update,ReplyKeyboardMarkup,ReplyKeyboardRemove
from telegram.ext import ApplicationBuilder, CommandHandler, ContextTypes, ConversationHandler, MessageHandler,filters
import requests
import os


BACKEND_URI=os.getenv('BACKEND_URI')
PROMETHEUS_URI=os.getenv('PROMETHEUS_URI')
TOKEN =os.getenv('API_TOKEN')

quantity=15
CHOOSING,GET_METRIC_TIME,GET_METRICS,GET_PLAYER_NAME,GET_PLAYER_PHONE,GET_PLAYER_EMAIL,GET_MATCH_ID,GET_PLAYER_MATCH_ID, CREATE_CHOOSING, CREATE_MATCH,GET_MATCH_NAME,GET_MATCH_DATE,DONE,GET_MATCH_PLACE,VIEW_CHOOSING = range(quantity)


logger = logging.getLogger(__name__)

async def hello(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(f'Hello {update.effective_user.first_name}')

async def view(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    reply_keyboard = [["Partido", "Metricas"]]
    await update.message.reply_text("Respecto a que quieres saber, partido o metricas?",
        reply_markup=ReplyKeyboardMarkup(
            reply_keyboard, one_time_keyboard=True, input_field_placeholder="Partido o Metricas?"
        ),
    )
    return VIEW_CHOOSING

async def create(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    reply_keyboard = [["Partido", "Jugador"]]
    await update.message.reply_text(
        "Que tipo de creacion queres realizar?"
        "Partido o Jugador en un partido?",
        reply_markup=ReplyKeyboardMarkup(
            reply_keyboard, one_time_keyboard=True, input_field_placeholder="Partido o Jugador?"
        ),
    )
    return CREATE_CHOOSING


async def init_conversation(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    reply_keyboard = [["Crear", "Ver"]]

    await update.message.reply_text(
        """
Hola, soy el bot dedicado a responder tus consultas.
Envia /cancel en cualquier momento para cancelar la conversacion.
Quieres crear o ver?
        """,
        reply_markup=ReplyKeyboardMarkup(
            reply_keyboard, one_time_keyboard=True, input_field_placeholder="Crear o Ver?"
        ),
    )

    return CHOOSING

async def create_match(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(
        """
Vamos a crear un partido, primero que descripcion tiene?
        """,
    )

    return GET_MATCH_NAME

async def create_player(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    context.user_data["player"] = {}
    await update.message.reply_text(
"""
Vamos a asignar un jugador a un partido, primero que indiqueme el identificador de su partido:
""",
    )

    return GET_PLAYER_MATCH_ID
async def get_player_match_id(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    print(context.user_data)
    context.user_data["player"]["match_id"] = update.message.text

    await update.message.reply_text(
"""
Cual es el nombre del jugador?
""",
    )
    return GET_PLAYER_NAME
async def get_player_name(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    context.user_data["player"]["name"] = update.message.text

    await update.message.reply_text(
"""
Cual es el telefono del jugador?
""",
    )
    return GET_PLAYER_PHONE
async def get_player_phone(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    context.user_data["player"]["phone"] = update.message.text

    await update.message.reply_text(
"""
Cual es el email del jugador?
""",
    )
    return GET_PLAYER_EMAIL
async def get_player_email(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    context.user_data["player"]["email"] = update.message.text
    matchId = context.user_data["player"]["match_id"]
    r = requests.post(f"{BACKEND_URI}/matches/{matchId}/player", json = context.user_data["player"])
    if r.status_code != 204:
        await update.message.reply_text("Datos invalidos")
        return
    await update.message.reply_text("Jugador Asignado")

async def get_match_name(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    match_data = {}
    match_data["description"] = update.message.text
    context.user_data["match"] = match_data
    await update.message.reply_text(
        """
En que lugar?
        """,
    )

    return GET_MATCH_PLACE

async def get_match_place(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    context.user_data["match"]["place"] = update.message.text
    await update.message.reply_text(
        """
En que fecha se va a realizar? formato: "dd/mm/yyyy"
        """,
    )

    return GET_MATCH_DATE


async def get_match_date(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    date = update.message.text
    try:
        date = datetime.datetime.strptime(date,"%d/%m/%Y")
    except ValueError as err:
        await update.message.reply_text("Ingrese una fecha valida")
        print(err)
        return GET_MATCH_DATE

    match_data = context.user_data["match"] | {
        "format": 10,
        "maxPlayers": 26,
        "date" : date.isoformat()+"Z"
    }
    print(match_data)
    
    r = requests.post(f"{BACKEND_URI}/matches", json = match_data)
    isValid = 200 <= r.status_code < 300
    if not isValid :
        await update.message.reply_text("algo salio mal")
        print(r.text)
        return
    await update.message.reply_text("Partido creado, id asignado:")
    await update.message.reply_text(r.text , reply_markup=ReplyKeyboardRemove())
    return ConversationHandler.END

async def done(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    print(context)
    await update.message.reply_text(
        """
Hasta Luego!
        """,
    )

    return

async def cancel(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    """Cancels and ends the conversation."""
    user = update.message.from_user
    logger.info("User %s canceled the conversation.", user.first_name)
    await update.message.reply_text(
        "Chau, espero haber sido de utilidad!", reply_markup=ReplyKeyboardRemove()
    )

    return ConversationHandler.END

async def view_match(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    await update.message.reply_text(
        """
Vamos a ver un partido, para esto necesito que me indique su codigo indentificador:
        """, reply_markup=ReplyKeyboardRemove()
    )

    return GET_MATCH_ID

async def get_match_id(update: Update, context: ContextTypes.DEFAULT_TYPE) -> None:
    id = update.message.text

    r = requests.get(f"{BACKEND_URI}/matches/{id}")
    if r.status_code != 200:
        await update.message.reply_text("id invalido, ingreselo nuevamente:",  reply_markup=ReplyKeyboardRemove())
        return GET_MATCH_ID
    #TODO: all
    res = r.json()
    finalizo = "si"
    finished = res.get("finished")
    if finished != "true":
        finalizo = "no"
    description = res.get("description")
    fecha = res.get("date")[0:10]
    lugar = res.get("place")
    def getName(player):
        return player["name"]
    titulares = list(map(getName,res.get("startingPlayers")))
    suplentes = list(map(getName,res.get("substitutePlayer")))
    partido = f"""
Descripcion: {description}
Finalizo: {finalizo}
Fecha: {fecha}
Lugar: {lugar}
Titulares: {titulares}
Suplentes: {suplentes}
    """
    await update.message.reply_text(partido, reply_markup=ReplyKeyboardRemove())
    return ConversationHandler.END


async def view_metrics(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    reply_keyboard = [["Partidos Creados"],["Jugadores Anotados"]]

    await update.message.reply_text(
        "Metricas de?",
        reply_markup=ReplyKeyboardMarkup(
            reply_keyboard, one_time_keyboard=True
        ),
    )

    return GET_METRICS
async def get_metrics(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    reply_keyboard = [["15m","30m","45m"],["1h","2h","3h"]]
    metrica = update.message.text
    metric = "created_matches"
    if metrica != "Partidos Creados":
        metric = "annotated_users"
    context.user_data["metric_desc"] = metrica
    context.user_data["metric"] = metric
    await update.message.reply_text(
        "Desde hace?",
        reply_markup=ReplyKeyboardMarkup(
            reply_keyboard, one_time_keyboard=True
        ),
    )
    return GET_METRIC_TIME

# TODO: get metric change to api not more prometheus
async def get_metric_time(update: Update, context: ContextTypes.DEFAULT_TYPE) -> int:
    metric = context.user_data["metric"]
    metricTime = update.message.text
    r = requests.get(f"{PROMETHEUS_URI}/api/v1/query?query={metric}&step={metricTime}")
    if r.status_code != 200:
        print(r.text)
        await update.message.reply_text("algo salio mal")
    count = r.json().get("data").get("result")[0].get("value")[1]
    metric_desc = context.user_data["metric_desc"]
    await update.message.reply_text(f"{metric_desc} : {count}" , reply_markup=ReplyKeyboardRemove())
    return ConversationHandler.END
    
app = ApplicationBuilder().token(TOKEN).build()

app.add_handler(CommandHandler("hello", hello))

# create match
conv_handler = ConversationHandler(
    entry_points=[CommandHandler("empezar", init_conversation)],
    states={
        CHOOSING: [
            MessageHandler(filters.Regex("^(Crear)$"),create),
            MessageHandler(filters.Regex("^(Ver)$"),view),
            ],
        VIEW_CHOOSING: [
            MessageHandler(filters.Regex("^(Partido)$"),view_match),
            MessageHandler(filters.Regex("^(Metricas)$"),view_metrics),
        ],
        GET_METRICS: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_metrics
            )
        ],
        GET_METRIC_TIME: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_metric_time,
            )
        ],
        CREATE_CHOOSING: [
            MessageHandler(filters.Regex("^(Partido)$"),create_match),
            MessageHandler(filters.Regex("^(Jugador)$"),create_player),
        ],
        GET_PLAYER_MATCH_ID: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_player_match_id,
            ),
        ],
        GET_PLAYER_NAME: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_player_name,
            ),
        ],
        GET_PLAYER_PHONE: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_player_phone,
            ),
        ],
        GET_PLAYER_EMAIL: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_player_email,
            ),
        ],
        GET_MATCH_ID: [            
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_match_id,
            ),
        ],
        GET_MATCH_NAME: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_match_name,
            ),
        ],
        GET_MATCH_PLACE: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_match_place,
            ),
        ],
        GET_MATCH_DATE: [
            MessageHandler(
                filters.TEXT & ~(filters.COMMAND | filters.Regex("^cancel$")),
                get_match_date,
            ),
        ]
    },
    fallbacks=[
        CommandHandler("cancel", cancel)
        ],
)

app.add_handler(conv_handler)


app.run_polling()